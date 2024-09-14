package repository

import "kioureki-vol4-backend/internal/domain/entity"


type OrganizationRepository interface {
	CreateOrganization(org entity.Organization) (*entity.Organization, error)
	FindByID(id uint) (*entity.Organization, error)
}
