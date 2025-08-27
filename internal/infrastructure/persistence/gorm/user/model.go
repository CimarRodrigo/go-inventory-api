package usergorm

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email         string    `gorm:"uniqueIndex;not null"`
	Name          string    `gorm:"not null"`
	Password      string    `gorm:"not null"`
	Status        string    `gorm:"not null"`
	CreatedDate   time.Time `gorm:"autoCreateTime"`
	UpdatedDate   time.Time `gorm:"autoUpdateTime"`
	LastLoginDate time.Time
}
