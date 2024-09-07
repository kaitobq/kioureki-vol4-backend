package controller

import "go-template/internal/usecase"

type UserController struct {
	uc usecase.UserUsecase
}

func NewUserController(uc usecase.UserUsecase) *UserController {
	return &UserController{
		uc: uc,
	}
}
