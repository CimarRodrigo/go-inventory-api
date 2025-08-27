package userio

import "github.com/CimarRodrigo/go-inventory-api/internal/domain/shared"

type UpdateInput struct {
	Email    string
	Name     string
	Password string
	Status   shared.Status
}
