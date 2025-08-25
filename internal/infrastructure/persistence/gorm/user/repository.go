package usergorm

import (
	userdomain "github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
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

func (repo *Repository) Create(user *userdomain.User) (*userdomain.User, error) {

	var newUser = FromDomain(user)

	if err := repo.db.Create(newUser).Error; err != nil {
		return nil, err
	}
	return ToDomain(newUser), nil
}

func (repo *Repository) GetByID(id uuid.UUID) (*userdomain.User, error) {
	var user User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return ToDomain(&user), nil
}

func (repo *Repository) GetAll() ([]*userdomain.User, error) {
	var users []User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return ToDomainList(users), nil
}
