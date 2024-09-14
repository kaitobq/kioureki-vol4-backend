package repository

import "kioureki-vol4-backend/internal/domain/entity"


type UserOrganizationMembershipRepository interface {
	CreateMembership(userID string, orgID string) error
	FindByUserID(userID string) (*[]entity.UserOrganizationMembership, error)
}
