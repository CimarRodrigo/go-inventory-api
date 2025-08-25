package userusecase

import (
	userdomain "github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
	"github.com/CimarRodrigo/go-inventory-api/pkg/logger"
	"github.com/google/uuid"
)

type DeleteUsecase struct {
	userRepo userdomain.Repository
	logger   logger.Logger
}

func NewDeleteUsecase(userRepo userdomain.Repository, logger logger.Logger) *DeleteUsecase {
	return &DeleteUsecase{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (uc *DeleteUsecase) Delete(id uuid.UUID) error {
	uc.logger.Debug("Starting user deletion process", "id", id)

	if err := uc.userRepo.Delete(id); err != nil {
		uc.logger.Error("User deletion failed", "id", id, "error", err)
		return err
	}

	uc.logger.Info("User deleted successfully", "id", id)
	return nil
}
