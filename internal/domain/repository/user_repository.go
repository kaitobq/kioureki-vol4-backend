package repository

import "kioureki-vol4-backend/internal/domain/entity"


type UserRepository interface {
	CheckDuplicateEmail(email string) (bool, error)
	CreateUser(user entity.User) error
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
	GetUserByEmail(email string) (*entity.User, error)
}
