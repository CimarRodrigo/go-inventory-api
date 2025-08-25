package userio

import "github.com/google/uuid"

type ListOneOutput struct {
	ID    uuid.UUID
	Email string
	Name  string
}
