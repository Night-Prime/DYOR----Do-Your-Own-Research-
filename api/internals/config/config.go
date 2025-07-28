package config

import (
	"errors"
	"os"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"github.com/Night-Prime/DYOR----Do-Your-Own-Research-.git/api/internals/database"
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

	VANTAGE_URL string
	VANTAGE_KEY string
	VANTAGE_FUNCTION string
	VANTAGE_FUNCTION_TOP string

	SecretKey string
	AIEndpoint string
	AIKey string
	AI_MODEL string

	REALTIME bool
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

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return nil, errors.New("SECRET_KEY not set in environment variables")
	}

	vantageURL := os.Getenv("VANTAGE_URL")
	if vantageURL == "" {
		return nil, errors.New("VANTAGE_URL not set in environment variables")
	}

	vantageKey := os.Getenv("VANTAGE_KEY")
	if vantageKey == "" {
		return nil, errors.New("VANTAGE_KEY not set in environment variables")
	}

	vantageFunction := os.Getenv("VANTAGE_FUNCTION")
	if vantageFunction == "" {
		return nil, errors.New("VANTAGE_FUNCTION not set in environment variables")
	}

	vantageFunctionTop := os.Getenv("VANTAGE_FUNCTION_TOP")
	if vantageFunctionTop == "" {
		return nil, errors.New("VANTAGE_FUNCTION_TOP not set in environment variables")
	}
	aiEndpoint := os.Getenv("AI_ENDPOINT")
	if aiEndpoint == "" {
		return nil, errors.New("AI_ENDPOINT not set in the environment variables")
	}

	aiKey := os.Getenv("AI_KEY")
	if aiKey == "" {
		return nil, errors.New("AI_KEY not set in the environment variables")
	}

	aiModel := os.Getenv("AI_MODEL")
	if aiModel == "" {
		return nil, errors.New("AI_MODEL not set in the environment variables")
	}
	
	realTime := os.Getenv("ALLOW_REALTIME")
	if realTime == "" {
		return nil, errors.New("ALLOW_REALTIME not set in the environment variables")
	}
	// converting the realTime string to boolean:
	value, err := strconv.ParseBool(realTime)
	if err != nil {
		return nil, errors.New("Error Converting Value")
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
		SecretKey: secretKey,
		VANTAGE_URL: vantageURL,
		VANTAGE_KEY: vantageKey,
		VANTAGE_FUNCTION: vantageFunction,
		VANTAGE_FUNCTION_TOP: vantageFunctionTop,
		AIEndpoint:aiEndpoint,
		AIKey: aiKey,
		AI_MODEL: aiModel,
		REALTIME: value,
	}

	if err := database.InitializeWithRetry(cfg.ConnStr); err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	return cfg, nil
}

func Get() *Config {
	if cfg == nil {
		panic("Failed to load config")
	}
	return cfg
}