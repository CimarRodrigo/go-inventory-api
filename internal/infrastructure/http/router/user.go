package router

import (
	userusecase "github.com/CimarRodrigo/go-inventory-api/internal/application/user/usecase"
	userdomain "github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
	userhandler "github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/http/handler/user"
	usergorm "github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/persistence/gorm/user"
	"github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/security/bcrypt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(r *gin.RouterGroup, db *gorm.DB) {

	// deps
	hasher := bcrypt.NewBcryptHasher()

	// repository
	userRepo := usergorm.NewRepository(db)

	// service
	userService := userdomain.NewService()

	// usecase
	userCreateUsecase := userusecase.NewCreateUseCase(userService, userRepo, hasher)
	userListUsecase := userusecase.NewListUseCase(userRepo)
	userUpdateUsecase := userusecase.NewUpdateUsecase(userRepo, hasher, userService)

	// handler
	userCreateHandler := userhandler.NewHandler(userCreateUsecase)
	userListHandler := userhandler.NewListHandler(userListUsecase)
	userUpdateHandler := userhandler.NewUpdateHandler(userUpdateUsecase)

	rg := r.Group("/users")
	{
		rg.POST("/", userCreateHandler.Create)
		rg.GET("/:id", userListHandler.GetByID)
		rg.GET("/", userListHandler.GetAll)
		rg.PATCH("/:id", userUpdateHandler.Update)
	}
}
