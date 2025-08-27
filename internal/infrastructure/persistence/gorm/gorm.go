package gorm

import (
	"fmt"
	"log"

	"github.com/CimarRodrigo/go-inventory-api/internal/config"
	usergorm "github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/persistence/gorm/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg config.Database) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to database")

	if err := db.AutoMigrate(
		&usergorm.User{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Tables migrated successfully")

	return db
}
