package roledomain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ValidateName(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return errors.New("role name is required")
	}
	if len(name) < 3 {
		return errors.New("role name must be at least 3 characters long")
	}
	return nil
}

func (s *Service) CreateRole(name, description string) (*Role, error) {
	name = strings.TrimSpace(name)
	description = strings.TrimSpace(description)

	if err := s.ValidateName(name); err != nil {
		return nil, err
	}

	now := time.Now()
	newRole := &Role{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		CreatedDate: now,
		UpdatedDate: now,
	}

	return newRole, nil
}

func (s *Service) UpdateRole(role *Role, name, description string) (*Role, error) {
	name = strings.TrimSpace(name)
	description = strings.TrimSpace(description)

	if err := s.ValidateName(name); err != nil {
		return nil, err
	}

	role.Name = name
	role.Description = description
	role.UpdatedDate = time.Now()

	return role, nil
}
