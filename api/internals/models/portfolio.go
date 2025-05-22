package models

import (
	"time"

	"fmt"
	"github.com/google/uuid"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/errors"
		"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/config"
)

type Portfolio struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Name		string	`gorm:"type:varchar(100);not null" json:"name"`
	TotalValue float64   `gorm:"type:float;not null" json:"total_value"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
	Assets   []Asset   `gorm:"foreignKey:PortfolioID" json:"assets,omitempty"`
}

func SavePortfolioToDB(p *Portfolio) error {
	p.ID = uuid.New()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	db := config.LoadDB()
	go AutoMigrate()

	var existingPortfolio Portfolio
	if err := db.Where("user_id = ? ", p.UserID).First(&existingPortfolio).Error; err == nil {
		return fmt.Errorf("Portfolio with %v already exists for User", p.UserID)
	}

	if err := db.Create(p).Error; err != nil {
		return fmt.Errorf("error saving portfolio to database: %v", err)
	}

	return nil
}

func DeletePortfolio(portfolioID string) error {
	db := config.LoadDB()

	if portfolioID == "" {
		return fmt.Errorf("Portfolio ID is required for deletion")
	}

	var portfolio Portfolio
	if err := db.First(&portfolio, "id = ?", portfolioID).Error; err != nil {
		return fmt.Errorf("Portfolio does not exist")
	}

	if err := db.Delete(&portfolio).Error; err != nil {
		return fmt.Errorf("Error Deleting Portfolio: %v", err)
	}

	return nil
}
