package router

import (
	"github.com/CimarRodrigo/go-inventory-api/internal/config"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, cfg *config.Config) {
	v1 := r.Group("/api/v1")
	{
		SetupHealthRoutes(v1, cfg)
	}
}
