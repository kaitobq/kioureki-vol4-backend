package repository

import "kioureki-vol4-backend/internal/domain/entity"

type UserOrganizationInvitationRepository interface {
	CreateInvitation(invitation entity.UserOrganizationInvitation, invitedBy string) error
}