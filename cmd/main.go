package main

import (
	"log"

	"github.com/CimarRodrigo/go-inventory-api/internal/config"
	"github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/http/router"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()

	gin.SetMode(cfg.Server.Mode)

	r := gin.Default()

	router.SetupRoutes(r, cfg)

	port := ":" + cfg.Server.Port

	log.Println("Starting server on", port)

	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
