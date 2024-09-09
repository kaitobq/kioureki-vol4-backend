package usecase

import "go-template/internal/usecase/response"

type UserUsecase interface {
	CreateUser(name, email, password string) (*response.SignUpResponse, error)
}
