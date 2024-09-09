package db

import (
	"fmt"
	"go-template/internal/domain/entity"
	"go-template/internal/domain/repository"
	"go-template/pkg/database"

	"golang.org/x/crypto/bcrypt"
)

type userRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user entity.User) (*int64, error) {
	query := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`
	
	result, err := r.db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *userRepository) HashPassword(password string) (string, error) {
    if password == "" || len(password) == 0 {
        return "", fmt.Errorf("password is required")
    }

    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        return "", err
    }

    return string(bytes), nil
}
