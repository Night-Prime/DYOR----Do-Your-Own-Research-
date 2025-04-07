package models

import (
	"time"

	"github.com/google/uuid"
)

type AssetType string

const (
	AssetTypeStock  AssetType = "stock"
	AssetTypeBond   AssetType = "bond"
	AssetTypeCrypto AssetType = "crypto"
)

type Asset struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	PortfolioID   uuid.UUID `json:"portfolio_id" gorm:"type:uuid;not null"`
	Symbol        string    `json:"symbol" gorm:"not null"`
	Name          string    `json:"name"`
	Type          AssetType `json:"type" gorm:"not null"`
	Quantity      float64   `json:"quantity"`
	PurchasePrice float64   `json:"purchase_price"`
	CurrentPrice  float64   `json:"current_price"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
