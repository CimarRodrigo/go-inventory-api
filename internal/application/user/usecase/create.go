package userusecase

import (
	"log"

	userio "github.com/CimarRodrigo/go-inventory-api/internal/application/user/io"
	"github.com/CimarRodrigo/go-inventory-api/internal/domain/shared"
	userdomain "github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
)

type CreateUseCase struct {
	userService    *userdomain.Service
	userRepo       userdomain.Repository
	passwordHasher shared.PasswordHasher
}

func NewCreateUseCase(userService *userdomain.Service, userRepo userdomain.Repository, passwordHasher shared.PasswordHasher) *CreateUseCase {
	return &CreateUseCase{
		userService:    userService,
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
	}
}

func (uc *CreateUseCase) Create(req *userio.CreateInput) (*userio.CreateOutput, error) {
	if err := uc.userService.ValidateEmail(req.Email); err != nil {
		log.Printf("Failed to validate email %s: %v", req.Email, err)
		return nil, err
	}

	if err := uc.userService.ValidatePassword(req.Password); err != nil {
		log.Printf("Failed to validate password: %v", err)
		return nil, err
	}

	hashedPassword, err := uc.passwordHasher.Hash(req.Password)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return nil, err
	}

	newUser, err := uc.userService.CreateUser(req.Email, hashedPassword, req.Name, shared.Active)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return nil, err
	}

	createdUser, err := uc.userRepo.Create(newUser)
	if err != nil {
		log.Printf("Failed to persist user: %v", err)
		return nil, err
	}

	response := &userio.CreateOutput{
		Email: createdUser.Email,
		Name:  createdUser.Name,
	}

	log.Printf("User created: %v", response)

	return response, nil

}
