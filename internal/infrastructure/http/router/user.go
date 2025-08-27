package router

import (
	"log"

	userusecase "github.com/CimarRodrigo/go-inventory-api/internal/application/user/usecase"
	userdomain "github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
	userhandler "github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/http/handler/user"
	"github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/logger"
	usergorm "github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/persistence/gorm/user"
	"github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/security/bcrypt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoutes(r *gin.RouterGroup, db *gorm.DB) {

	// deps
	hasher := bcrypt.NewBcryptHasher()
	logger, err := logger.NewZapLogger(true) // development = true
	if err != nil {
		log.Fatal(err)
	}

	// repository
	userRepo := usergorm.NewRepository(db)

	// service
	userService := userdomain.NewService()

	// usecase
	userCreateUsecase := userusecase.NewCreateUseCase(userService, userRepo, hasher, logger)
	userListUsecase := userusecase.NewListUseCase(userRepo, logger)
	userUpdateUsecase := userusecase.NewUpdateUsecase(userRepo, hasher, userService, logger)
	userDeleteUsecase := userusecase.NewDeleteUsecase(userRepo, logger)

	// handler
	userCreateHandler := userhandler.NewHandler(userCreateUsecase)
	userListHandler := userhandler.NewListHandler(userListUsecase)
	userUpdateHandler := userhandler.NewUpdateHandler(userUpdateUsecase)
	userDeleteHandler := userhandler.NewDeleteHandler(userDeleteUsecase)

	rg := r.Group("/users")
	{
		rg.POST("/", userCreateHandler.Create)
		rg.GET("/:id", userListHandler.GetByID)
		rg.GET("/", userListHandler.GetAll)
		rg.PATCH("/:id", userUpdateHandler.Update)
		rg.DELETE("/:id", userDeleteHandler.Delete)
	}
}
