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
	userCreateUsecase := appuser.NewCreateUseCase(userService, userRepo, hasher)
	userListUsecase := appuser.NewListUseCase(userRepo)

	// handler
	userCreateHandler := handleruser.NewHandler(userCreateUsecase)
	userListHandler := handleruser.NewListHandler(userListUsecase)

	rg := r.Group("/users")
	{
		rg.POST("/", userCreateHandler.Create)
		rg.GET("/:id", userListHandler.GetByID)
	}
}
