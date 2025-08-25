package userusecase

import (
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
		return nil, err
	}

	return ToListOneOutput(user), nil
}

func (uc *ListUseCase) GetAll() ([]*userio.ListOneOutput, error) {
	users, err := uc.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return ToListOneOutputList(users), nil
}
