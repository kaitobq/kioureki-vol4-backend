package usecase

import "go-template/internal/domain/repository"

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}
