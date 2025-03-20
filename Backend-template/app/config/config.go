package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_SSL      string
	POSTGRES_TIMEZONE string
	ENV               string
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvOrThrow(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}

func NewConfig() *Config {
	// Load .env file

	dotenverr := godotenv.Load()
	if dotenverr != nil {
		log.Printf("Warning: .env file not found or error loading. Using system environment variables.")
	}
	return &Config{
		POSTGRES_USER:     getEnvOrThrow("POSTGRES_USER"),
		POSTGRES_PASSWORD: getEnvOrThrow("POSTGRES_PASSWORD"),
		POSTGRES_DB:       getEnvOrThrow("POSTGRES_DB"),
		POSTGRES_HOST:     getEnvOrDefault("POSTGRES_HOST", "localhost"),
		POSTGRES_PORT:     getEnvOrDefault("POSTGRES_PORT", "5432"),
		POSTGRES_SSL:      getEnvOrDefault("POSTGRES_SSL", "disable"),
		POSTGRES_TIMEZONE: getEnvOrDefault("POSTGRES_TIMEZONE", "Asia/Bangkok"),
		ENV:               getEnvOrDefault("ENV", "dev"),
	}
}
