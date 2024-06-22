package models

import "github.com/jinzhu/gorm"

type CategoryOption struct {
	gorm.Model
	CategoryID uint `json:"category_id"`
	Value	  string `gorm:"size:255;not null" json:"value"`
}