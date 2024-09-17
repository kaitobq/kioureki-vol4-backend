package repository

import "kioureki-vol4-backend/internal/domain/entity"


type UserOrganizationMembershipRepository interface {
	CreateMembership(userID, orgID, role string) error
	FindByUserID(userID string) (*[]entity.UserOrganizationMembership, error)
}
