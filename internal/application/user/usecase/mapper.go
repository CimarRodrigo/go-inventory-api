package userusecase

import (
	userio "github.com/CimarRodrigo/go-inventory-api/internal/application/user/io"
	userdomain "github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
)

func ToListOneOutput(u *userdomain.User) *userio.ListOneOutput {
	return &userio.ListOneOutput{
		ID:            u.ID,
		Email:         u.Email,
		Name:          u.Name,
		CreatedDate:   u.CreatedDate,
		UpdatedDate:   u.UpdatedDate,
		LastLoginDate: u.LastLoginDate,
		Status:        u.Status,
	}
}

func ToListOneOutputList(users []*userdomain.User) []*userio.ListOneOutput {
	result := make([]*userio.ListOneOutput, len(users))
	for i, u := range users {
		result[i] = ToListOneOutput(u)
	}
	return result
}
