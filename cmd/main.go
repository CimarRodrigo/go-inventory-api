package main

import (
	"log"

	"github.com/CimarRodrigo/go-inventory-api/internal/config"
	"github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/http/router"
	"github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/persistence/gorm"
	"github.com/gin-gonic/gin"
)

func main() {

	//Initial config
	cfg := config.LoadConfig()
	gin.SetMode(cfg.Server.Mode)
	r := gin.Default()

	//Database
	db := gorm.InitDB(cfg.Database)

	//Router
	router.SetupRoutes(r, cfg, db)

	//Init server
	port := ":" + cfg.Server.Port
	log.Println("Starting server on", port)
	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
