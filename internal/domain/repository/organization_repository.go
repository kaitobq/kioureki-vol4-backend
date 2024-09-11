package repository

import "go-template/internal/domain/entity"

type OrganizationRepository interface {
	CreateOrganization(org entity.Organization) (*entity.Organization, error)
}
