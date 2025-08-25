package router

import (
	appuser "github.com/CimarRodrigo/go-inventory-api/internal/application/user/usecase"
	domainuser "github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
	handleruser "github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/http/handler/user"
	gormuser "github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/persistence/gorm/user"
	"github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/security/bcrypt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(r *gin.RouterGroup, db *gorm.DB) {

	// deps
	hasher := bcrypt.NewBcryptHasher()

	// repository
	userRepo := gormuser.NewRepository(db)

	// service
	userService := domainuser.NewService()

	// usecase
	userUsecase := appuser.NewCreateUseCase(userService, userRepo, hasher)

	// handler
	userHandler := handleruser.NewHandler(userUsecase)

	rg := r.Group("/users")
	{
		rg.POST("/", userHandler.Create)
	}
}
