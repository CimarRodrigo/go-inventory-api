package userusecase

import (
	userio "github.com/CimarRodrigo/go-inventory-api/internal/application/user/io"
	"github.com/CimarRodrigo/go-inventory-api/internal/domain/shared"
	userdomain "github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
	"github.com/CimarRodrigo/go-inventory-api/pkg/logger"
)

type CreateUseCase struct {
	userService    *userdomain.Service
	userRepo       userdomain.Repository
	passwordHasher shared.PasswordHasher
	logger         logger.Logger
}

func NewCreateUseCase(userService *userdomain.Service, userRepo userdomain.Repository, passwordHasher shared.PasswordHasher, logger logger.Logger) *CreateUseCase {
	return &CreateUseCase{
		userService:    userService,
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		logger:         logger,
	}
}

func (uc *CreateUseCase) Create(req *userio.CreateInput) (*userio.CreateOutput, error) {
	uc.logger.Debug("Starting user creation process", "email", req.Email)

	if err := uc.userService.ValidateEmail(req.Email); err != nil {
		uc.logger.Warn("Email validation failed", "email", req.Email, "error", err)
		return nil, err
	}

	if err := uc.userService.ValidatePassword(req.Password); err != nil {
		uc.logger.Warn("Password validation failed", "error", err)
		return nil, err
	}

	hashedPassword, err := uc.passwordHasher.Hash(req.Password)
	if err != nil {
		uc.logger.Error("Password hashing failed", "error", err)
		return nil, err
	}

	newUser, err := uc.userService.CreateUser(req.Email, hashedPassword, req.Name, shared.Active)
	if err != nil {
		uc.logger.Error("User domain creation failed", "email", req.Email, "error", err)
		return nil, err
	}

	createdUser, err := uc.userRepo.Create(newUser)
	if err != nil {
		uc.logger.Error("User persistence failed", "email", req.Email, "error", err)
		return nil, err
	}

	response := &userio.CreateOutput{
		Email: createdUser.Email,
		Name:  createdUser.Name,
	}

	uc.logger.Info("User created successfully", "email", response.Email)

	return response, nil
}
