package entity

import "time"

type UserOrganizationMembership struct {
	ID             uint
	UserID         uint
	OrganizationID uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
