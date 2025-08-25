package userhandler

import (
	userusecase "github.com/CimarRodrigo/go-inventory-api/internal/application/user/usecase"
	userdto "github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/http/dto/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateHandler struct {
	Usecase *userusecase.UpdateUsecase
}

func NewUpdateHandler(usecase *userusecase.UpdateUsecase) *UpdateHandler {
	return &UpdateHandler{
		Usecase: usecase,
	}
}

func (h *UpdateHandler) Update(c *gin.Context) {
	var req userdto.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)

	input := ToUpdateInput(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user ID"})
		return
	}
	if err := h.Usecase.UpdateInfo(input, id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User updated successfully"})
}
