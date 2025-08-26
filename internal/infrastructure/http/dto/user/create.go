package userdto

type RegisterRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
