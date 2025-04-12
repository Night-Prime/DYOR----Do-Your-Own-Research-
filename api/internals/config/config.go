package config

import (
	"errors"
	"os"
	
	"github.com/joho/godotenv"
)

var cfg *Config

type Config struct {
	Port    string
	ConnStr string

	StockAPI_URL string
	StockAPI_Key  string
	StockHostname string

	CryptoAPI_URL string
	CryptoAPI_Key  string
	CryptoHostname string
}

func Load() (*Config, error) {
	var err error
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

	stockAPI_URL := os.Getenv("STOCK_URL")
	if stockAPI_URL == "" {
		return nil, errors.New("STOCK_URL not set in environment variables")
	}

	stockAPI_Key := os.Getenv("STOCK_API_KEY")
	if stockAPI_Key == "" {
		return nil, errors.New("STOCK_API_KEY not set in environment variables")
	}

	stockHostname := os.Getenv("STOCK_API_HOST")
	if stockHostname == "" {
		return nil, errors.New("STOCK_HOSTNAME not set in environment variables")
	}

	cryptoAPI_URL := os.Getenv("CRYPTO_URL")
	if cryptoAPI_URL == "" {
		return nil, errors.New("CRYPTO_URL not set in environmental variables")
	}

	cryptoAPI_Key := os.Getenv("CRYPTO_API_KEY")
	if cryptoAPI_Key == "" {
		return nil, errors.New("CRYPTO_API_KEY not set in environmental variables")
	}

	cryptoHostname := os.Getenv("CRYPTO_API_HOST")
	if cryptoHostname == "" {
		return nil, errors.New("CRYPTO_API_HOST not set in environmental variables")
	}

	cfg =  &Config{
		Port: port,
		ConnStr: connStr,
		StockAPI_URL: stockAPI_URL,
		StockAPI_Key: stockAPI_Key,
		StockHostname: stockHostname,
		CryptoAPI_Key: cryptoAPI_Key,
		CryptoAPI_URL: cryptoAPI_URL,
		CryptoHostname: cryptoHostname,
	}

	return cfg, err
}

func Get() *Config {
	if cfg == nil {
		panic("Failed to load config")
	}
	return cfg
}