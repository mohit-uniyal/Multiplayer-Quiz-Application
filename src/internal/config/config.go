package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT                    string
	MONGO_CONNECTION_STRING string
	MONGO_DATABASE_NAME     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".secrets/.env")
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	config := Config{
		PORT:                    os.Getenv("PORT"),
		MONGO_CONNECTION_STRING: os.Getenv("MONGO_CONNECTION_STRING"),
		MONGO_DATABASE_NAME:     os.Getenv("MONGO_DATABASE_NAME"),
	}

	return &config, nil
}
