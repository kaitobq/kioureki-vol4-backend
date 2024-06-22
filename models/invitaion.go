package models

import "github.com/jinzhu/gorm"

type Invitation struct {
	gorm.Model
	OrganizationID uint
	UserID         uint
}
