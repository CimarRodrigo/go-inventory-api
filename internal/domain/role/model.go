package roledomain

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedDate time.Time
	UpdatedDate time.Time
}
