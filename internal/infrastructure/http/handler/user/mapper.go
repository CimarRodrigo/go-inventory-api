package user

import (
	userio "github.com/CimarRodrigo/go-inventory-api/internal/application/user/io"
	userdto "github.com/CimarRodrigo/go-inventory-api/internal/infrastructure/http/dto/user"
)

func ToInput(req *userdto.CreateRequest) *userio.CreateInput {
	return &userio.CreateInput{
		Email:    req.Email,
		Name:     req.Name,
		Password: req.Password,
	}
}
