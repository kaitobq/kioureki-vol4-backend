package repository

import "go-template/internal/domain/entity"

type UserOrganizationMembershipRepository interface {
	CreateMembership(userID, orgID uint) error
	FindByUserID(userID uint) (*[]entity.UserOrganizationMembership, error)
}
