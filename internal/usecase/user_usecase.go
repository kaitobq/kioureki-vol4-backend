package usecase

import (
	"go-template/internal/domain/entity"
	"go-template/internal/domain/repository"
	"go-template/internal/domain/service"
	"go-template/internal/usecase/response"
)

type userUsecase struct {
	repo         repository.UserRepository
	TokenService service.TokenService
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (uc *userUsecase) CreateUser(name, email, password string) (*response.SignUpResponse, error) {
	hashedPassword, err := uc.repo.HashPassword(password)
	if err != nil {
		return nil, err
	}
	
	user := entity.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	id, err := uc.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	token, err := uc.TokenService.GenerateTokenFromID(uint(*id))
	if err != nil {
		return nil, err
	}

	exp, err := uc.TokenService.ExtractExpFromToken(token)
	if err != nil {
		return nil, err
	}

	return response.NewSignUpResponse(token, exp)
}
