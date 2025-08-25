package userusecase

import (
	userio "github.com/CimarRodrigo/go-inventory-api/internal/application/user/io"
	userdomain "github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
	"github.com/CimarRodrigo/go-inventory-api/pkg/logger"
	"github.com/google/uuid"
)

type ListUseCase struct {
	userRepo userdomain.Repository
	logger   logger.Logger
}

func NewListUseCase(userRepo userdomain.Repository, logger logger.Logger) *ListUseCase {
	return &ListUseCase{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (uc ListUseCase) GetByID(id uuid.UUID) (*userio.ListOneOutput, error) {
	uc.logger.Debug("Retrieving user by ID", "id", id)

	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		uc.logger.Error("Failed to retrieve user", "id", id, "error", err)
		return nil, err
	}

	return ToListOneOutput(user), nil
}

func (uc ListUseCase) GetAll() ([]*userio.ListOneOutput, error) {
	uc.logger.Debug("Retrieving all users")

	users, err := uc.userRepo.GetAll()
	if err != nil {
		uc.logger.Error("Failed to retrieve users", "error", err)
		return nil, err
	}

	uc.logger.Debug("Users retrieved successfully", "count", len(users))
	return ToListOneOutputList(users), nil
}
