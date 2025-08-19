package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig
	App    AppConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type AppConfig struct {
	Name        string
	Version     string
	Environment string
}

func LoadConfig() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Mode: getEnv("GIN_MODE", "debug"), // debug, release, test
		},
		App: AppConfig{
			Name:        getEnv("APP_NAME", "My API"),
			Version:     getEnv("APP_VERSION", "1.0.0"),
			Environment: getEnv("APP_ENV", "development"), // development, staging, production
		},
	}
}

func getEnv(key, defaultValue string) string {

	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
