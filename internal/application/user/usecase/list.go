package userusecase

import (
	"log"

	userio "github.com/CimarRodrigo/go-inventory-api/internal/application/user/io"
	userdomain "github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
	"github.com/google/uuid"
)

type ListUseCase struct {
	userRepo userdomain.Repository
}

func NewListUseCase(userRepo userdomain.Repository) *ListUseCase {
	return &ListUseCase{
		userRepo: userRepo,
	}
}

func (uc *ListUseCase) GetByID(id uuid.UUID) (*userio.ListOneOutput, error) {
	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		log.Printf("Failed to get user by ID %s: %v", id.String(), err)
		return nil, err
	}

	log.Printf("Retrieved user by ID %s", id.String())
	return ToListOneOutput(user), nil
}

func (uc *ListUseCase) GetAll() ([]*userio.ListOneOutput, error) {
	users, err := uc.userRepo.GetAll()
	if err != nil {
		log.Printf("Failed to get all users: %v", err)
		return nil, err
	}

	log.Printf("Retrieved %d users", len(users))
	return ToListOneOutputList(users), nil
}
