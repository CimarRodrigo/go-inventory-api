package user

import "github.com/google/uuid"

type Repository interface {
	Create(user *User) error
	GetByID(id uuid.UUID) (*User, error)
	Update(user *User) error
	GetByEmail(email string) (*User, error)
	Delete(id uuid.UUID) error
}
