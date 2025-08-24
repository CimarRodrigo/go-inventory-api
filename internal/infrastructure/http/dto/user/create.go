package userdto

type CreateRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
