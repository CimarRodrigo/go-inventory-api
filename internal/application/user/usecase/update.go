package userusecase

import (
	"log"

	userio "github.com/CimarRodrigo/go-inventory-api/internal/application/user/io"
	"github.com/CimarRodrigo/go-inventory-api/internal/domain/shared"
	userdomain "github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
	"github.com/google/uuid"
)

type UpdateUsecase struct {
	userRepo       userdomain.Repository
	passwordHasher shared.PasswordHasher
	userService    *userdomain.Service
}

func NewUpdateUsecase(userRepo userdomain.Repository, passwordHasher shared.PasswordHasher, userService *userdomain.Service) *UpdateUsecase {
	return &UpdateUsecase{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		userService:    userService,
	}
}

func (uc *UpdateUsecase) UpdateInfo(user *userio.UpdateInput, id uuid.UUID) error {
	if user.Password != "" {

		if err := uc.userService.ValidatePassword(user.Password); err != nil {
			log.Printf("Failed to validate password: %v", err)
			return err
		}

		hashedPassword, err := uc.passwordHasher.Hash(user.Password)
		if err != nil {
			log.Printf("Failed to hash password: %v", err)
			return err
		}
		user.Password = hashedPassword
	}
	log.Printf("Updating user with ID %s", id.String())

	updatedUser, err := uc.userService.UpdateUser(user.Email, user.Password, user.Name)
	if err != nil {
		return err
	}

	return uc.userRepo.Update(updatedUser, id)
}
