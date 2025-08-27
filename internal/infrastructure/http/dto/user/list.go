package userdto

import (
	"time"

	"github.com/google/uuid"
)

type ListOneResponse struct {
	ID            uuid.UUID `json:"id"`
	Email         string    `json:"email"`
	Name          string    `json:"name"`
	Status        string    `json:"status"`
	CreatedDate   time.Time `json:"created_date"`
	UpdatedDate   time.Time `json:"updated_date"`
	LastLoginDate time.Time `json:"last_login_date"`
}
