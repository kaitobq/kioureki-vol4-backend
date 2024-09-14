package entity

import "time"

type UserOrganizationMembership struct {
	ID             uint
	UserID         string
	OrganizationID string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
