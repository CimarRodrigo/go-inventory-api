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

	return &userio.ListOneOutput{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func (uc *ListUseCase) GetAll() ([]*userio.ListOneOutput, error) {
	users, err := uc.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []*userio.ListOneOutput
	for _, user := range users {
		result = append(result, &userio.ListOneOutput{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		})
	}

	return result, nil
}
