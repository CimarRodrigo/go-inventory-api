package userdto

import "github.com/google/uuid"

type ListOneResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}
