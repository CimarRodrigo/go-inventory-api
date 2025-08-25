package userusecase

import (
	userio "github.com/CimarRodrigo/go-inventory-api/internal/application/user/io"
	"github.com/CimarRodrigo/go-inventory-api/internal/domain/shared"
	userdomain "github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
	"github.com/CimarRodrigo/go-inventory-api/pkg/logger"
	"github.com/google/uuid"
)

type UpdateUsecase struct {
	userRepo       userdomain.Repository
	passwordHasher shared.PasswordHasher
	userService    *userdomain.Service
	logger         logger.Logger
}

func NewUpdateUsecase(userRepo userdomain.Repository, passwordHasher shared.PasswordHasher, userService *userdomain.Service, logger logger.Logger) *UpdateUsecase {
	return &UpdateUsecase{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		userService:    userService,
		logger:         logger,
	}
}

func (uc UpdateUsecase) UpdateInfo(user *userio.UpdateInput, id uuid.UUID) error {
	uc.logger.Debug("Starting user update process", "id", id)

	if user.Password != "" {
		uc.logger.Debug("Processing password update", "id", id)

		if err := uc.userService.ValidatePassword(user.Password); err != nil {
			uc.logger.Warn("Password validation failed", "id", id, "error", err)
			return err
		}

		hashedPassword, err := uc.passwordHasher.Hash(user.Password)
		if err != nil {
			uc.logger.Error("Password hashing failed", "id", id, "error", err)
			return err
		}
		user.Password = hashedPassword
	}

	updatedUser, err := uc.userService.UpdateUser(user.Email, user.Password, user.Name)
	if err != nil {
		uc.logger.Error("User domain update failed", "id", id, "error", err)
		return err
	}

	if err := uc.userRepo.Update(updatedUser, id); err != nil {
		uc.logger.Error("User persistence failed", "id", id, "error", err)
		return err
	}

	uc.logger.Info("User updated successfully", "id", id, "email", updatedUser.Email)
	return nil
}
