package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	OrganizationID uint `json:"organization_id"`
	Name 		 string `gorm:"size:255;not null" json:"name"`
	Options		 []CategoryOption `json:"options" gorm:"foreignkey:CategoryID"`
}


func (c *Category) Create(name string, organizationId uint, options []string) error {
	c.Name = name
	c.OrganizationID = organizationId

	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&c).Error; err != nil {
			return err
		}
		for _, option := range options {
			categoryOption := CategoryOption{CategoryID: c.ID, Value: option}
			if err := tx.Create(&categoryOption).Error; err != nil {
				return err
			}
		}
		return nil
	})

	return err
}