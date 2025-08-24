package io

type CreateInput struct {
	Email    string
	Name     string
	Password string
}

type CreateOutput struct {
	Email string
	Name  string
}
