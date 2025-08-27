package userhandler

import (
	"net/http"

	userusecase "github.com/CimarRodrigo/go-inventory-api/internal/application/user/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteHandler struct {
	Usecase *userusecase.DeleteUsecase
}

func NewDeleteHandler(usecase *userusecase.DeleteUsecase) *DeleteHandler {
	return &DeleteHandler{
		Usecase: usecase,
	}
}

func (h *DeleteHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	if err := h.Usecase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
