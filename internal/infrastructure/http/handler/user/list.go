package userhandler

import (
	"net/http"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	user, err := h.Usecase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToListOneResponse(user))
}

func (h *ListHandler) GetAll(c *gin.Context) {
	users, err := h.Usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ToListOneResponseList(users))
}
