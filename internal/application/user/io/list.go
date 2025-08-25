package userio

import (
	"time"

	"github.com/CimarRodrigo/go-inventory-api/internal/domain/shared"
	"github.com/google/uuid"
)

type ListOneOutput struct {
	ID            uuid.UUID
	Email         string
	Name          string
	CreatedDate   time.Time
	UpdatedDate   time.Time
	LastLoginDate time.Time
	Status        shared.Status
}
