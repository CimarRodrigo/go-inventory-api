package usergorm

import (
	"github.com/CimarRodrigo/go-inventory-api/internal/domain/shared"
	domainuser "github.com/CimarRodrigo/go-inventory-api/internal/domain/user"
)

func FromDomain(user *domainuser.User) *User {
	return &User{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Status:   string(shared.Active),
	}
}

func ToDomain(user *User) *domainuser.User {
	return &domainuser.User{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Status:   shared.Status(user.Status),
	}
}

func ToDomainList(users []User) []*domainuser.User {
	var result []*domainuser.User
	for _, user := range users {
		result = append(result, ToDomain(&user))
	}
	return result
}
