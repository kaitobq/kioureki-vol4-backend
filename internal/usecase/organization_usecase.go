package usecase

import (
	"fmt"
	"kioureki-vol4-backend/internal/domain/entity"
	"kioureki-vol4-backend/internal/domain/repository"
	"kioureki-vol4-backend/internal/usecase/response"
)

type organizationUsecase struct {
	repo repository.OrganizationRepository
	membershipRepo repository.UserOrganizationMembershipRepository
}

func NewOrganizationUsecase(repo repository.OrganizationRepository, membershipRepo repository.UserOrganizationMembershipRepository) OrganizationUsecase {
	return &organizationUsecase{
		repo: repo,
		membershipRepo: membershipRepo,
	}
}

func (uc *organizationUsecase) CreateOrganization(name string, founderID uint) (*response.CreateOrganizationResponse, error) {
	org := entity.Organization{
		Name: name,
	}

	memberships, err := uc.membershipRepo.FindByUserID(founderID)
	if err != nil {
		return nil, err
	}

	if len(*memberships) != 0 {
		for _, membership := range *memberships {
			org, err := uc.repo.FindByID(membership.OrganizationID)
			if err != nil {
				return nil, err
			}

			if org.Name == name {
				return nil, fmt.Errorf("cannot belong to multiple organizations with the same name [%s]", name)
			}
		}
	}

	organization, err := uc.repo.CreateOrganization(org)
	if err != nil {
		return nil, err
	}

	return response.NewCreateOrganizationResponse(organization.Name, organization.CreatedAt, organization.UpdatedAt)
}
