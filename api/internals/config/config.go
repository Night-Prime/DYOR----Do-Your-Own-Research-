package config

import (
	"errors"
	"os"
	
	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	ConnStr string
}

func Load() (*Config, error) {
	godotenv.Load()

	// grab port
	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("PORT not set in environment variables")
	}

	// grab DB connection string
	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		return nil, errors.New("DB_URL not set in environment variables")
	}

	return &Config{
		Port: port,
		ConnStr: connStr,
	}, nil
}