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
	// AssetTypeCrypto AssetType = "crypto"
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

// type CryptoData struct {
//     Image                   string  `json:"image"`
//     MarketCap               float64 `json:"market_cap"`
//     MarketCapRank           int     `json:"market_cap_rank"`
//     High24H                 float64 `json:"high_24h"`
//     Low24H                  float64 `json:"low_24h"`
//     PriceChange24H          float64 `json:"price_change_24h"`
//     PriceChangePercentage24H float64 `json:"price_change_percentage_24h"`
//     CirculatingSupply       float64 `json:"circulating_supply"`
//     TotalSupply             float64 `json:"total_supply"`
//     LastUpdated             string  `json:"last_updated"`
// }

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
                // Add other fields as needed...
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
    // CryptoData *CryptoData `json:"crypto_data,omitempty" gorm:"type:jsonb"`
    StockData  *StockData  `json:"stock_data,omitempty" gorm:"type:jsonb"`
    // BondData   *BondData   `json:"bond_data,omitempty" gorm:"type:jsonb"`
}

func (a *Asset) Validate() error {
    switch a.Type {
    // case AssetTypeCrypto:
    //     if a.CryptoData == nil {
    //         return errors.New("crypto_data is required for crypto assets")
    //     }
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