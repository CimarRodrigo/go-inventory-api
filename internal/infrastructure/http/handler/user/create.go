package userhandler

import (
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
	var req userdto.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	input := ToCreateInput(&req)

	user, err := h.Usecase.Create(input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, userdto.CreateResponse{
		Email: user.Email,
		Name:  user.Name,
	})
}
