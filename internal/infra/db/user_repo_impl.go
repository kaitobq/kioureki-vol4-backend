package db

import "go-template/pkg/database"

type userRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}
