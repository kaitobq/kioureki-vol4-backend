package usecase

import (
	"fmt"
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

// owner > admin > member
func (u *userOrganizationMembershipUsecase) DeleteMembership(executorID, deleteUserID, organizationID string) (*response.DeleteMembershipResponse, error) {
	executorRole, err := u.repo.GetRole(executorID, organizationID)
	if err != nil {
		return nil, err
	}

	deleteUserRole, err := u.repo.GetRole(deleteUserID, organizationID)
	if err != nil {
		return nil, err
	}

	isDeletable := false
	switch executorRole {
		case "owner":
			isDeletable = true
		case "admin":
			if deleteUserRole == "member" {
				isDeletable = true
			}
	}
	if !isDeletable {
		return nil, fmt.Errorf("permission denied")
	}

	err = u.repo.DeleteMembership(deleteUserID, organizationID)
	if err != nil {
		return nil, err
	}

	return response.NewDeleteMembershipResponse()
}
