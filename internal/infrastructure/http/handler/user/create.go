package userhandler

import (
	"net/http"

	userusecase "github.com/CimarRodrigo/go-inventory-api/internal/application/user/usecase"
	userdto "github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/http/dto/user"
	"github.com/gin-gonic/gin"
)

type CreateHandler struct {
	Usecase *userusecase.CreateUseCase
}

func NewHandler(usecase *userusecase.CreateUseCase) *CreateHandler {
	return &CreateHandler{
		Usecase: usecase,
	}
}

func (h *CreateHandler) Create(c *gin.Context) {
	var req userdto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input := ToCreateInput(&req)

	user, err := h.Usecase.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userdto.RegisterResponse{
		Email: user.Email,
		Name:  user.Name,
	})
}
