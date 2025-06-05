package models

import (
	"time"
	"fmt"

	"github.com/google/uuid"
    "github.com/lib/pq"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/errors"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/database"
)

type Portfolio struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID          uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Name		    string	`gorm:"type:varchar(100);not null" json:"name"`
	TotalValue      float64   `gorm:"type:float;not null" json:"total_value"`
	InvestmentGoals pq.StringArray `gorm:"type:text[]" json:"investment_goals"`
    AssetPreference pq.StringArray `gorm:"type:text[]" json:"asset_preference"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"index" json:"deleted_at"`
	Assets          []Asset   `gorm:"foreignKey:PortfolioID" json:"assets,omitempty"`
}

func SavePortfolioToDB(p *Portfolio) error {
	p.ID = uuid.New()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	db := database.GetDB()
	go AutoMigrate()

	var existingPortfolio Portfolio
	if err := db.Where("user_id = ? ", p.UserID).First(&existingPortfolio).Error; err == nil {
		return &errors.DatabaseError{ Message:"Portfolio with already exists for User", Err: err}
	}

	if err := db.Create(p).Error; err != nil {
		return &errors.DatabaseError{ Message: "error saving portfolio to database", Err: err}
	}

	return nil
}

func DeletePortfolio(portfolioID string) error {
	db := database.GetDB()

	if portfolioID == "" {
		return &errors.ValidationError{ Message:"Portfolio ID is required for deletion"}
	}

	var portfolio Portfolio
	if err := db.First(&portfolio, "id = ?", portfolioID).Error; err != nil {
		return &errors.DatabaseError{Message:"Portfolio does not exist", Err: err}
	}

	if err := db.Delete(&portfolio).Error; err != nil {
		return  &errors.DatabaseError{Message:"Error Deleting Portfolio", Err:err}
	}

	return nil
}


func UpdatePortfolio(p *Portfolio, assetsToAdd []*Asset, assetsToUpdate []*Asset, assetsToDelete []uuid.UUID) error {
    db := database.GetDB()
    p.UpdatedAt = time.Now()

    if p.ID == uuid.Nil {
        return &errors.ValidationError{Message: "Portfolio ID is required"}
    }

    // Verify portfolio exists
    var existingPortfolio Portfolio
    if err := db.Where("id = ? AND user_id = ?", p.ID, p.UserID).
        First(&existingPortfolio).Error; err != nil {
        return &errors.DatabaseError{
            Message: "Portfolio not found or doesn't belong to user", 
            Err: err,
        }
    }

    tx := db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // Update portfolio metadata
    updates := map[string]interface{}{
        "name":             p.Name,
        "total_value":      p.TotalValue,
        "updated_at":       p.UpdatedAt,
    }

    // Only update these fields if they're not nil
    if p.InvestmentGoals != nil {
        updates["investment_goals"] = p.InvestmentGoals
    }
    if p.AssetPreference != nil {
        updates["asset_preference"] = p.AssetPreference
    }

    if err := tx.Model(&Portfolio{}).Where("id = ?", p.ID).Updates(updates).Error; err != nil {
        tx.Rollback()
        return &errors.DatabaseError{Message: "Error updating portfolio", Err: err}
    }


    // Handle asset deletions first
    if len(assetsToDelete) > 0 {
        if err := tx.Where("portfolio_id = ? AND id IN ?", p.ID, assetsToDelete).
            Delete(&Asset{}).Error; err != nil {
            tx.Rollback()
            return &errors.DatabaseError{Message: "Error deleting assets", Err: err}
        }
    }

    // Handle asset updates
    for _, asset := range assetsToUpdate {
        if asset.ID == uuid.Nil {
            tx.Rollback()
            return &errors.ValidationError{Message: "Asset ID is required for updates"}
        }

        asset.UpdatedAt = time.Now()
        if err := tx.Model(&Asset{}).Where("id = ? AND portfolio_id = ?", asset.ID, p.ID).
            Updates(map[string]interface{}{
                "symbol":        asset.Symbol,
                "name":          asset.Name,
                "quantity":      asset.Quantity,
                "current_price": asset.CurrentPrice,
                "volume":        asset.Volume,
                "updated_at":    asset.UpdatedAt,
            }).Error; err != nil {
            tx.Rollback()
            return &errors.DatabaseError{Message: "Error updating asset", Err: err}
        }
    }

    // Handle new asset creations using your existing logic
    if len(assetsToAdd) > 0 {
        // Prepare assets
        var symbols []string
        for _, asset := range assetsToAdd {
            asset.AssetBase.ID = uuid.New()
            asset.AssetBase.CreatedAt = time.Now()
            asset.AssetBase.UpdatedAt = time.Now()
            asset.PortfolioID = p.ID
            symbols = append(symbols, asset.Symbol)
        }

        // Check for existing symbols
        var existingAssets []Asset
        if err := tx.Where("portfolio_id = ? AND symbol IN ?", p.ID, symbols).
            Find(&existingAssets).Error; err == nil && len(existingAssets) > 0 {
            tx.Rollback()
            existingSymbols := make([]string, len(existingAssets))
            for i, a := range existingAssets {
                existingSymbols[i] = a.Symbol
            }
            return fmt.Errorf("Assets with these symbols already exist: %v", existingSymbols)
        }

        if err := tx.CreateInBatches(assetsToAdd, 100).Error; err != nil {
            tx.Rollback()
            return &errors.DatabaseError{Message: "Error adding assets to portfolio", Err: err}
        }
    }

    return tx.Commit().Error
}
