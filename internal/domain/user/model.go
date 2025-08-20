package user

import (
	"time"

	"github.com/CimarRodrigo/go-inventory-api/internal/domain/shared"
	"github.com/google/uuid"
)

type User struct {
	Id            uuid.UUID
	Email         string
	Name          string
	Password      string
	Status        shared.Status
	CreatedDate   time.Time
	UpdatedDate   time.Time
	LastLoginDate time.Time
}
