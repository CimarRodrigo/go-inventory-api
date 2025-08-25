package userdomain

import "github.com/google/uuid"

type Repository interface {
	Create(user *User) (*User, error)
	GetByID(id uuid.UUID) (*User, error)
	GetAll() ([]*User, error)
	Update(user *User, id uuid.UUID) error
}
