package user

import (
	"errors"
	"regexp"

	"github.com/CimarRodrigo/go-inventory-api/internal/domain/shared"
	"github.com/google/uuid"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

func (s *Service) ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[.,/$()=%#]`).MatchString(password)

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}

	if !hasNumber {
		return errors.New("password must contain at least one number")
	}

	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func (s *Service) CreateUser(email string, password string, name string, status shared.Status) (*User, error) {
	newUser := &User{
		ID:       uuid.New(),
		Email:    email,
		Name:     name,
		Password: password,
		Status:   status,
	}

	return newUser, nil
}
