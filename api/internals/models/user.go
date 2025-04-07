package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	FirstName  string      `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName   string      `gorm:"type:varchar(100);not null" json:"last_name"`
	Avatar     *string     `gorm:"type:varchar(255)" json:"avatar"`
	Email      *string     `gorm:"type:varchar(100);unique" json:"email"`
	Password   string      `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt  time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  *time.Time  `gorm:"index" json:"deleted_at"`
	Portfolios []Portfolio `gorm:"foreignKey:UserID" json:"portfolios,omitempty"`
}
