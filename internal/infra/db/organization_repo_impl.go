package db

import (
	"go-template/internal/domain/entity"
	"go-template/internal/domain/repository"
	"go-template/pkg/database"
	"time"
)

type organizationRepository struct {
	db database.DB
}

func NewOrganizationRepository(db database.DB) repository.OrganizationRepository {
	return &organizationRepository{
		db: db,
	}
}

// 名前被りはOK, 正し、同じユーザが同じ名前の組織を作成することはできない
func (r *organizationRepository) CreateOrganization(org entity.Organization) (*entity.Organization, error) {
	query := `INSERT INTO organizations (name, created_at, updated_at) VALUES (?, ?, ?)`

	now := time.Now()
	org.CreatedAt = now
	org.UpdatedAt = now
	
	result, err := r.db.Exec(query, org.Name, org.CreatedAt, org.UpdatedAt)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	org.ID = uint(id)

	return &org, nil
}
