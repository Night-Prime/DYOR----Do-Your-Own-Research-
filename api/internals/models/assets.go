package models

import (
    "fmt"
	"time"
    "encoding/json"
    "database/sql/driver"

	"github.com/google/uuid"
    "github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/errors"
    "github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/database"
)

type AssetType string

const (
	AssetTypeStock  AssetType = "stock"
	AssetTypeCrypto AssetType = "crypto"
)

type AssetSymbol struct {
    Symbol string `json:"symbol"`
    Name   string `json:"name"`
}

type AssetRequest struct {
    Type        string        `json:"type"`
    Symbols     []AssetSymbol `json:"symbols"`
    PortfolioID uuid.UUID     `json:"portfolioID"`
}

type AssetUpdateResponse struct {
    Stocks []*Asset `json:"stocks,omitempty"`
    Crypto []*Asset `json:"crypto,omitempty"`
    Errors []string        `json:"errors,omitempty"`
}


// The models that gets sent to the DB
type AssetBase struct {
    ID            uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    PortfolioID   uuid.UUID `json:"portfolio_id" gorm:"type:uuid;not null"`
    Symbol        string    `json:"symbol""`
    Name          string    `json:"name"`
    Type          AssetType `json:"type" gorm:"not null;index"`
    Quantity      float64   `json:"quantity"`
    CurrentPrice  float64   `json:"current_price"`
    Volume        float64   `json:"volume"`
    CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CryptoData struct {
    Key                string `json:"key"`
    ID                 int    `json:"id"`
    Name               string `json:"name"`
    Symbol             string `json:"symbol"`
    Decimals           int    `json:"decimals"`
    Logo               string `json:"logo"`
    Rank               int    `json:"rank"`
    Price              float64 `json:"price"`
    MarketCap          float64 `json:"market_cap"`
    MarketCapDiluted   float64 `json:"market_cap_diluted"`
    Volume             float64 `json:"volume"`
    VolumeChange24H    float64 `json:"volume_change_24h"`
    Volume7D           float64 `json:"volume_7d"`
    Liquidity          float64 `json:"liquidity"`
    ATH                float64 `json:"ath"`
    ATL                float64 `json:"atl"`
    PriceChange1H      float64 `json:"price_change_1h"`
    PriceChange24H     float64 `json:"price_change_24h"`
    PriceChange7D      float64 `json:"price_change_7d"`
    PriceChange1M      float64 `json:"price_change_1m"`
    PriceChange1Y      float64 `json:"price_change_1y"`
}

func (s *CryptoData) Scan(value interface{}) error {
    if value == nil {
        return nil
    }
    bytes, ok := value.([]byte)
    if !ok {
        return fmt.Errorf("Failed to unmarshal Crypto Data: expected []byte, got %T", value)
    }
    return json.Unmarshal(bytes, s)
}

func (s CryptoData) Value() (driver.Value, error) {
    return json.Marshal(s)
}

type StockData struct {
    Symbol  string  `json:"symbol"`
    RegularMarketPrice struct {
        Raw float64 `json:"raw"`
    } `json:"regularMarketPrice"`
    MarketCap struct {
        Raw float64 `json:"raw"`
    } `json:"marketCap"`
    RegularMarketVolume struct {
        Raw float64 `json:"raw"`
    } `json:"regularMarketVolume"`
    Name              string  `json:"longName"`
    Exchange          string  `json:"exchange"`
    SharesOutstanding struct {
        Raw float64 `json:"raw"`
    } `json:"sharesOutstanding"`
    RegularMarketChange struct {
        Raw float64 `json:"raw"`
    } `json:"regularMarketChange"`
}

func (s *StockData) Scan(value interface{}) error {
    if value == nil {
        return nil
    }
    bytes, ok := value.([]byte)
    if !ok {
        return fmt.Errorf("Failed to unmarshal StockData: expected []byte, got %T", value)
    }
    return json.Unmarshal(bytes, s)
}

func (s StockData) Value() (driver.Value, error) {
    return json.Marshal(s)
}

type Asset struct {
    AssetBase
    CryptoData *CryptoData `json:"crypto_data,omitempty" gorm:"type:jsonb"`
    StockData  *StockData  `json:"stock_data,omitempty" gorm:"type:jsonb"`
}

// just for API Client Responses
type StockAPIResponse struct {
    Data struct {
        QuoteResponse struct {
            Result []StockData `json:"result"`
        } `json:"quoteResponse"`
    } `json:"data"`
}

type CryptoAPIResponse struct{
    Data      map[string]interface{} `json:"data"`
    DataArray []CryptoData `json:dataArray`
}



func SaveAssetToDB(assets []*Asset) error {
    fmt.Println("Saving to the DB")
    fmt.Println("--------------------------------------------- \n")

    db := database.GetDB()

    if len(assets) == 0 {
        return &errors.ValidationError{Message: "No assets provided"}
    }

    portfolioID := assets[0].PortfolioID
    var portfolio Portfolio

    if err := db.Where("id = ?", portfolioID).First(&portfolio).Error; err != nil {
        return &errors.DatabaseError{Message: "Portfolio does not exist", Err: err}
    }

    // Prepare assets
    var symbols []string
    for _, asset := range assets {
        asset.AssetBase.ID = uuid.New()
        asset.AssetBase.CreatedAt = time.Now()
        asset.AssetBase.UpdatedAt = time.Now()
        symbols = append(symbols, asset.Symbol)
    }

    tx := db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // Check for existing assets
    var existingAssets []Asset
    if err := tx.Where("portfolio_id = ? AND symbol IN ?", portfolioID, symbols).Find(&existingAssets).Error; err != nil {
        tx.Rollback()
        return &errors.DatabaseError{Message: "Error checking existing assets", Err: err}
    }

    // Create a map of existing symbols for quick lookup
    existingMap := make(map[string]bool)
    for _, a := range existingAssets {
        existingMap[a.Symbol] = true
    }

    // Separate new assets from existing ones
    var newAssets []*Asset
    var assetsToUpdate []*Asset

    for _, asset := range assets {
        if existingMap[asset.Symbol] {
            assetsToUpdate = append(assetsToUpdate, asset)
        } else {
            newAssets = append(newAssets, asset)
        }
    }

    // Create new assets
    if len(newAssets) > 0 {
        if err := tx.CreateInBatches(newAssets, 100).Error; err != nil {
            tx.Rollback()
            return &errors.DatabaseError{Message: "Error adding new assets to portfolio", Err: err}
        }
    }

    // Update existing assets
    for _, asset := range assetsToUpdate {
        if err := tx.Model(&Asset{}).
            Where("portfolio_id = ? AND symbol = ?", portfolioID, asset.Symbol).
            Updates(map[string]interface{}{
                "stock_data":   asset.StockData,   // Update these fields
                "crypto_data":  asset.CryptoData,  // as needed
                "updated_at":    time.Now(),
            }).Error; err != nil {
            tx.Rollback()
            return &errors.DatabaseError{Message: "Error updating existing assets", Err: err}
        }
    }

    return tx.Commit().Error
}

func UpdateAsset(asset *Asset) error {
    db := database.GetDB()
    asset.UpdatedAt = time.Now()
    
    if asset.ID == uuid.Nil {
        return &errors.ValidationError{Message: "Asset ID is required"}
    }

    var existingAsset Asset
    if err := db.First(&existingAsset, "id = ?", asset.ID).Error; err != nil {
        return &errors.DatabaseError{Message: "Asset not found", Err: err}
    }

    if err := db.Save(asset).Error; err != nil {
        return &errors.DatabaseError{Message: "Error updating asset", Err: err}
    }
    return nil
}

func DeleteAsset(assetID string) error {
    db := database.GetDB()

    if assetID == "" {
        return &errors.ValidationError{Message:"Asset ID is required for deletion"}
    }

    assetUUID, err := uuid.Parse(assetID)
    if err != nil {
        return &errors.ValidationError{Message:"Invalid asset ID format"}
    }

    var asset Asset

    if err := db.First(&asset, "id = ?", assetUUID).Error; err != nil {
        return &errors.DatabaseError{Message:"Error finding asset", Err:err}
    }

    if err := db.Delete(&asset).Error; err != nil {
        return  &errors.DatabaseError{Message:"Error deleting asset", Err:err}
    }

    return nil
}