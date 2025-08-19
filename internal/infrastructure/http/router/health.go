package router

import (
	"github.com/CimarRodrigo/go-inventory-api/internal/config"
	"github.com/gin-gonic/gin"
)

func SetupHealthRoutes(rg *gin.RouterGroup, cfg *config.Config) {
	health := rg.Group("/health")
	{
		health.GET("/", healthCheck(cfg))
	}
}

func healthCheck(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"app":     cfg.App.Name,
			"version": cfg.App.Version,
			"env":     cfg.App.Environment,
		})
	}
}
