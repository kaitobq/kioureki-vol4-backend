package models

import "github.com/jinzhu/gorm"

type Membership struct {
	gorm.Model
	OrganizationID uint
	UserID         uint
}
