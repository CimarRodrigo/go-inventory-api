package user

import "github.com/google/uuid"

type Repository interface {
	Create(user *User) (*User, error)
	GetByID(id uuid.UUID) (*User, error)
}
