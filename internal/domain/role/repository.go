package roledomain

type Repository interface {
	Create(role *Role) (*Role, error)
}
