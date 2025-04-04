package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Host     string
	Password string
	Db       int
}

func LoadConfig() (*Config, error) {
	db, err := strconv.Atoi(getEnvOrDefault("CACHE_DATABASE", "0"))
	if err != nil {
		return nil, fmt.Errorf("invalid database: %v", err)
	}

	return &Config{
		Host:     getEnvOrDefault("CACHE_URL", "localhost:6379"),
		Password: getEnvOrDefault("CACHE_PASSWORD", ""),
		Db:       db,
	}, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}
