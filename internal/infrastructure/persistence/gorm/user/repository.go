package usergorm

import (
	"github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) Create(user *user.User) (*user.User, error) {

	var newUser = FromDomain(user)

	if err := repo.db.Create(newUser).Error; err != nil {
		return nil, err
	}
	return ToDomain(newUser), nil
}

func (repo *Repository) GetByID(id uuid.UUID) (*user.User, error) {
	var user User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return ToDomain(&user), nil
}
