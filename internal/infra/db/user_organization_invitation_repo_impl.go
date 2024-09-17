package db

import (
	"kioureki-vol4-backend/internal/domain/entity"
	"kioureki-vol4-backend/internal/domain/repository"
	"kioureki-vol4-backend/pkg/database"
)

type userOrganizationInvitationRepository struct {
	db database.DB
}

func NewUserOrganizationInvitationRepository(db *database.DB) repository.UserOrganizationInvitationRepository {
	return &userOrganizationInvitationRepository{
		db: *db,
	}
}

func (r *userOrganizationInvitationRepository) CreateInvitation(invitation entity.UserOrganizationInvitation, invitedBy string) error {
	query := `INSERT INTO user_organization_invitations (id, organization_id, user_id, invited_by, created_at, updated_at) VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`

	_, err := r.db.Exec(query, invitation.ID, invitation.OrganizationID, invitation.UserID, invitedBy)
	if err != nil {
		return err
	}

	return nil
}
