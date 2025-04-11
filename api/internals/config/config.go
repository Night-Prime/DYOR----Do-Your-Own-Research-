package config

import (
	"errors"
	"os"
	
	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	ConnStr string

	StockAPI_URL string
	StockAPI_Key  string
	StockHostname string
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

	// grab stock API URL
	stockAPI_URL := os.Getenv("STOCK_URL")
	if stockAPI_URL == "" {
		return nil, errors.New("STOCK_URL not set in environment variables")
	}

	stockAPI_Key := os.Getenv("STOCK_API_KEY")
	if stockAPI_Key == "" {
		return nil, errors.New("STOCK_API_KEY not set in environment variables")
	}

	StockHostname := os.Getenv("STOCK_API_HOST")
	if StockHostname == "" {
		return nil, errors.New("STOCK_HOSTNAME not set in environment variables")
	}

	return &Config{
		Port: port,
		ConnStr: connStr,
		StockAPI_URL: stockAPI_URL,
		StockAPI_Key: stockAPI_Key,
		StockHostname: StockHostname,
	}, nil
}