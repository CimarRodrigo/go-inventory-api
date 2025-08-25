package userhandler

import (
	userusecase "github.com/CimarRodrigo/go-inventory-api/internal/application/user/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ListHandler struct {
	Usecase *userusecase.ListUseCase
}

func NewListHandler(usecase *userusecase.ListUseCase) *ListHandler {
	return &ListHandler{
		Usecase: usecase,
	}
}

func (h *ListHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid UUID"})
		return
	}
	user, err := h.Usecase.GetByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, ToListOneResponse(user))
}

func (h *ListHandler) GetAll(c *gin.Context) {
	users, err := h.Usecase.GetAll()
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, ToListOneResponseList(users))
}
