package db

import (
	"go-template/internal/domain/entity"
	"go-template/internal/domain/repository"
	"go-template/pkg/database"
	"time"
)

type userOrganizationMembershipRepository struct {
	db database.DB
}

func NewUserOrganizationMembershipRepository(db database.DB) repository.UserOrganizationMembershipRepository {
	return &userOrganizationMembershipRepository{
		db: db,
	}
}

func (r *userOrganizationMembershipRepository) CreateMembership(userID, orgID uint) error {
	query := `INSERT INTO user_organization_memberships (user_id, organization_id, created_at, updated_at) VALUES (?, ?, ?, ?)`

	now := time.Now()
	_, err := r.db.Exec(query, userID, orgID, now, now)
	if err != nil {
		return err
	}

	return nil
}

func (r *userOrganizationMembershipRepository) FindByUserID(userID uint) (*[]entity.UserOrganizationMembership, error) {
	query := `SELECT * FROM user_organization_memberships WHERE user_id = ?`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var memberships []entity.UserOrganizationMembership
	for rows.Next() {
		var membership entity.UserOrganizationMembership
		if err := rows.Scan(&membership.ID, &membership.UserID, &membership.OrganizationID, &membership.CreatedAt, &membership.UpdatedAt); err != nil {
			return nil, err
		}
		memberships = append(memberships, membership)
	}

	return &memberships, nil
}
