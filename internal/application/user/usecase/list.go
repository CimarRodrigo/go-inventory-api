package userusecase

import (
	userio "github.com/CimarRodrigo/go-inventory-api/internal/application/user/io"
	"github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
	"github.com/google/uuid"
)

type ListUseCase struct {
	userRepo user.Repository
}

func NewListUseCase(userRepo user.Repository) *ListUseCase {
	return &ListUseCase{
		userRepo: userRepo,
	}
}

func (uc *ListUseCase) GetByID(id uuid.UUID) (*userio.ListOneOutput, error) {
	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &userio.ListOneOutput{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
