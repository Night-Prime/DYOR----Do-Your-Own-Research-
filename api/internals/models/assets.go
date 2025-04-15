package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type AssetType string

const (
	AssetTypeStock  AssetType = "stock"
	// AssetTypeBond   AssetType = "bond"
	AssetTypeCrypto AssetType = "crypto"
)

type AssetBase struct {
    ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
    PortfolioID   uuid.UUID `json:"portfolio_id" gorm:"type:uuid;not null"`
    Symbol        string    `json:"symbol" gorm:"not null"`
    Name          string    `json:"name"`
    Type          AssetType `json:"type" gorm:"not null;index"`
    Quantity      float64   `json:"quantity"`
    CurrentPrice  float64   `json:"current_price"`
    Volume        float64   `json:"volume"`
    CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type CryptoData struct {
    Data      map[string]interface{} `json:"data"`
    DataArray []struct {
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
    } `json:"dataArray"`
}

type StockData struct {
    Status int `json:"status" gorm:"-"`
    Data   struct {
        QuoteResponse struct {
            Result []struct {
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
            } `json:"result"`
        } `json:"quoteResponse"`
    } `json:"data"`
}

// type BondData struct {
//     YieldToMaturity   float64 `json:"yield_to_maturity"`
//     MaturityDate      string  `json:"maturity_date"`
//     CouponRate        float64 `json:"coupon_rate"`
//     CreditRating      string  `json:"credit_rating"`
// }

type Asset struct {
    AssetBase
    CryptoData *CryptoData `json:"crypto_data,omitempty" gorm:"type:jsonb"`
    StockData  *StockData  `json:"stock_data,omitempty" gorm:"type:jsonb"`
    // BondData   *BondData   `json:"bond_data,omitempty" gorm:"type:jsonb"`
}

func (a *Asset) Validate() error {
    switch a.Type {
    case AssetTypeCrypto:
        if a.CryptoData == nil {
            return errors.New("crypto_data is required for crypto assets")
        }
    case AssetTypeStock:
        if a.StockData == nil {
            return errors.New("stock_data is required for stock assets")
        }
    // case AssetTypeBond:
    //     if a.BondData == nil {
    //         return errors.New("bond_data is required for bond assets")
    //     }
    default:
        return errors.New("invalid asset type")
    }
    return nil
}