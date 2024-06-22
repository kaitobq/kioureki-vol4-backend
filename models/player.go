package models

import "github.com/jinzhu/gorm"

type Player struct {
	gorm.Model
	Name         string `gorm:"size:255;not null" json:"name"`
	OrganizationID uint `json:"organization_id"`
}
