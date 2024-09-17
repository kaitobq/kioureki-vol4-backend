package usecase

import (
	"kioureki-vol4-backend/internal/domain/repository"
	"kioureki-vol4-backend/internal/usecase/response"
)

type userOrganizationMembershipUsecase struct {
	repo repository.UserOrganizationMembershipRepository
}

func NewUserOrganizationMembershipUsecase(repo repository.UserOrganizationMembershipRepository) UserOrganizationMembershipUsecase {
	return &userOrganizationMembershipUsecase{
		repo: repo,
	}
}

func (u *userOrganizationMembershipUsecase) CreateMembership(userID string, organizationID string) (*response.CreateMembershipResponse, error) {
	err := u.repo.CreateMembership(userID, organizationID, "member")
	if err != nil {
		return nil, err
	}

	return response.NewCreateMembershipResponse()
}
