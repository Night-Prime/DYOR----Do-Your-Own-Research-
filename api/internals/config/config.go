package config

import (
	"errors"
	"os"
	
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func Load() (*Config, error) {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("PORT not set in environment variables")
	}

	return &Config{
		Port: port,
	}, nil
}