package usecase

import (
	"github.com/CimarRodrigo/go-inventory-api/internal/application/user/io"
	"github.com/CimarRodrigo/go-inventory-api/internal/domain/shared"
	"github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
)

type CreateUseCase struct {
	userService    *user.Service
	userRepo       user.Repository
	passwordHasher shared.PasswordHasher
}

func NewCreateUseCase(userService *user.Service, userRepo user.Repository, passwordHasher shared.PasswordHasher) *CreateUseCase {
	return &CreateUseCase{
		userService:    userService,
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
	}
}

func (uc *CreateUseCase) Create(req *io.CreateInput) (*io.CreateOutput, error) {

	if err := uc.userService.ValidateEmail(req.Email); err != nil {
		return nil, err
	}

	if err := uc.userService.ValidatePassword(req.Password); err != nil {
		return nil, err
	}

	hashedPassword, err := uc.passwordHasher.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	newUser, err := uc.userService.CreateUser(req.Email, hashedPassword, req.Name, shared.Active)
	if err != nil {
		return nil, err
	}

	createdUser, err := uc.userRepo.Create(newUser)
	if err != nil {
		return nil, err
	}

	response := &io.CreateOutput{
		Email: createdUser.Email,
		Name:  createdUser.Name,
	}

	return response, nil

}
