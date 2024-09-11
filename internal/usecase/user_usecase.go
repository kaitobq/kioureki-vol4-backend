package usecase

import (
	"fmt"
	"go-template/internal/domain/entity"
	"go-template/internal/domain/repository"
	"go-template/internal/domain/service"
	"go-template/internal/usecase/response"
)

type userUsecase struct {
	repo         repository.UserRepository
	TokenService service.TokenService
}

func NewUserUsecase(repo repository.UserRepository, tokenService service.TokenService) UserUsecase {
	return &userUsecase{
		repo: repo,
		TokenService: tokenService,
	}
}

func (uc *userUsecase) CreateUser(name, email, password string) (*response.SignUpResponse, error) {
	exists, err := uc.repo.CheckDuplicateEmail(email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("email already exists")
	}

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

func (uc *userUsecase) SignIn(email, password string) (*response.SignInResponse, error) {
	user, err := uc.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	err = uc.repo.ComparePassword(user.Password, password)
	if err != nil {
		return nil, err
	}
	
	token, err := uc.TokenService.GenerateTokenFromID(user.ID)
	if err != nil {
		return nil, err
	}

	exp, err := uc.TokenService.ExtractExpFromToken(token)
	if err != nil {
		return nil, err
	}

	return response.NewSignInResponse(token, exp)
}
