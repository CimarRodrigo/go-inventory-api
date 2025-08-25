package router

import (
	"github.com/CimarRodrigo/go-inventory-api/internal/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Bootstrap(r *gin.Engine, cfg *config.Config, db *gorm.DB) {

	// routes
	v1 := r.Group("/api/v1")
	{
		SetupHealthRoutes(v1, cfg)
		SetupUserRoutes(v1, db)
	}

}
