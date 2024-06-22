package models

import "github.com/jinzhu/gorm"

type Organization struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name"`
}
