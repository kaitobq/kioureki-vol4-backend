package repository

import "go-template/internal/domain/entity"

type UserRepository interface {
	CreateUser(user entity.User) (*int64, error)
	HashPassword(password string) (string, error)
}
