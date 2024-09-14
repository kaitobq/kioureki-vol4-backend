package repository

import "kioureki-vol4-backend/internal/domain/entity"


type UserOrganizationMembershipRepository interface {
	CreateMembership(userID, orgID uint) error
	FindByUserID(userID uint) (*[]entity.UserOrganizationMembership, error)
}
