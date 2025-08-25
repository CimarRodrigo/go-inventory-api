package userhandler

import (
	userio "github.com/CimarRodrigo/go-inventory-api/internal/application/user/io"
	userdto "github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/http/dto/user"
)

func ToCreateInput(req *userdto.CreateRequest) *userio.CreateInput {
	return &userio.CreateInput{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}
}

func ToListOneResponse(user *userio.ListOneOutput) *userdto.ListOneResponse {
	return &userdto.ListOneResponse{
		ID:            user.ID,
		Email:         user.Email,
		Name:          user.Name,
		CreatedDate:   user.CreatedDate,
		UpdatedDate:   user.UpdatedDate,
		LastLoginDate: user.LastLoginDate,
		Status:        string(user.Status),
	}
}

func ToListOneResponseList(users []*userio.ListOneOutput) []*userdto.ListOneResponse {
	result := make([]*userdto.ListOneResponse, len(users))
	for i, u := range users {
		result[i] = ToListOneResponse(u)
	}
	return result
}
