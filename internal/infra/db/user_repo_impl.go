package db

import (
	"fmt"
	"kioureki-vol4-backend/internal/domain/entity"
	"kioureki-vol4-backend/internal/domain/repository"
	"kioureki-vol4-backend/pkg/database"

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

func (r *userRepository) CheckDuplicateEmail(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)`

	var exists bool
	err := r.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *userRepository) CreateUser(user entity.User) error {
	query := `INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`
	
	_, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
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

func (r *userRepository) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User, error) {
	query := `SELECT id, name, email, password FROM users WHERE email = ?`

	var user entity.User
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
