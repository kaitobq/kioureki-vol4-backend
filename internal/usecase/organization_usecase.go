package usecase

import (
	"go-template/internal/domain/entity"
	"go-template/internal/domain/repository"
	"go-template/internal/usecase/response"
)

type organizationUsecase struct {
	repo repository.OrganizationRepository
}

func NewOrganizationUsecase() OrganizationUsecase {
	return &organizationUsecase{}
}

func (uc *organizationUsecase) CreateOrganization(name string) (*response.CreateOrganizationResponse, error) {
	org := entity.Organization{
		Name: name,
	}

	organization, err := uc.repo.CreateOrganization(org)
	if err != nil {
		return nil, err
	}

	return response.NewCreateOrganizationResponse(organization.Name, organization.CreatedAt, organization.UpdatedAt)
}
