package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".secrets/.env")
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	config := Config{
		PORT: os.Getenv("PORT"),
	}

	return &config, nil
}
