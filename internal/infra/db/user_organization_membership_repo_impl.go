package db

import (
	"kioureki-vol4-backend/internal/domain/entity"
	"kioureki-vol4-backend/internal/domain/repository"
	"kioureki-vol4-backend/pkg/database"
	"time"
)

type userOrganizationMembershipRepository struct {
	db database.DB
}

func NewUserOrganizationMembershipRepository(db *database.DB) repository.UserOrganizationMembershipRepository {
	return &userOrganizationMembershipRepository{
		db: *db,
	}
}

func (r *userOrganizationMembershipRepository) CreateMembership(userID, orgID, role string) error {
	query := `INSERT INTO user_organization_memberships (user_id, organization_id, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`

	now := time.Now()
	_, err := r.db.Exec(query, userID, orgID, role, now, now)
	if err != nil {
		return err
	}

	return nil
}

func (r *userOrganizationMembershipRepository) FindByUserID(userID string) (*[]entity.UserOrganizationMembership, error) {
	query := `SELECT * FROM user_organization_memberships WHERE user_id = ?`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var memberships []entity.UserOrganizationMembership
	for rows.Next() {
		var membership entity.UserOrganizationMembership
		if err := rows.Scan(&membership.ID, &membership.UserID, &membership.OrganizationID, &membership.Role, &membership.CreatedAt, &membership.UpdatedAt); err != nil {
			return nil, err
		}
		memberships = append(memberships, membership)
	}

	return &memberships, nil
}
