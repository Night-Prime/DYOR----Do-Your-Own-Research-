package models

import (
	"time"

	"github.com/google/uuid"
)

type Portfolio struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Name		string	`gorm:"type:varchar(100);not null" json:"name"`
	TotalValue float64   `gorm:"type:float;not null" json:"total_value"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
	Assets   []Asset   `gorm:"foreignKey:PortfolioID" json:"assets,omitempty"`
}