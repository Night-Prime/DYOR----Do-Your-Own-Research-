package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db          *gorm.DB
	maxRetries  = 5
	retryDelay  = 5 * time.Second
)

// InitializeWithRetry sets up the database connection with retry logic
func InitializeWithRetry(connStr string) error {
	var err error
	
	for attempt := 1; attempt <= maxRetries; attempt++ {
		db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
			PrepareStmt: true,
		})
		
		if err == nil {
			sqlDB, err := db.DB()
			if err != nil {
				return fmt.Errorf("Failed to get database pool: %w", err)
			}

			// Configure connection pool
			sqlDB.SetMaxIdleConns(10)
			sqlDB.SetMaxOpenConns(100)
			sqlDB.SetConnMaxLifetime(2 * time.Hour)
			sqlDB.SetConnMaxIdleTime(30 * time.Minute)

			// Test connection
			if err := sqlDB.Ping(); err == nil {
				return nil
			}
		}

		log.Printf("Database connection attempt %d/%d failed: %v", attempt, maxRetries, err)
		if attempt < maxRetries {
			log.Printf("Retrying in %v...", retryDelay)
			time.Sleep(retryDelay)
		}
	}

	return fmt.Errorf("Failed to connect to database after %d attempts: %w", maxRetries, err)
}

// GetDB returns the initialized database instance
func GetDB() *gorm.DB {
	return db
}

// HealthCheck verifies the database connection
func HealthCheck() error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("Failed to get database pool: %w", err)
	}
	return sqlDB.Ping()
}