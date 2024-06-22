package models

import (
	"backend/utils/token"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct{
	gorm.Model
	Username string `gorm:"size:255;not null;" json:"username"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

func IsEmailExist(email string) bool {
	var user User
	err := DB.Where("email = ?", email).First(&user).Error
	return err == nil
}


func (u *User) BeforeSave() error{
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	u.Username = strings.ToLower(u.Username)

	return nil
}

func (u User) Save() (User, error){
	err := DB.Create(&u).Error
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (u User) PrepareOutput() User{
	u.Password = ""
	return u
}

func GenerateToken(email string, password string) (string, error){
	var user User

	err := DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", err
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func UpdateUsername(userId uint, username string, password string) (string, error) {
	var user User

	if err := DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	user.Username = username
	if err := DB.Model(&user).Update("username", username).Error; err != nil {
		return "", err
	}

	return user.Username, nil
}
