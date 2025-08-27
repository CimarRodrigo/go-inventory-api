package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   Server
	App      App
	Database Database
}

func LoadConfig() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	return &Config{
		Server:   LoadServerConfig(),
		App:      LoadAppConfig(),
		Database: LoadDatabaseConfig(),
	}
}
