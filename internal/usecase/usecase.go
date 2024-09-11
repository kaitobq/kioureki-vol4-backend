package usecase

import "go-template/internal/usecase/response"

type UserUsecase interface {
	CreateUser(name, email, password string) (*response.SignUpResponse, error)
	SignIn(email, password string) (*response.SignInResponse, error)
}

type OrganizationUsecase interface {
	CreateOrganization(name string) (*response.CreateOrganizationResponse, error)
}
