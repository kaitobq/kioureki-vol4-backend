package usecase

import "kioureki-vol4-backend/internal/usecase/response"


type UserUsecase interface {
	CreateUser(name, email, password string) (*response.SignUpResponse, error)
	SignIn(email, password string) (*response.SignInResponse, error)
}

type OrganizationUsecase interface {
	CreateOrganization(name string, founderID string) (*response.CreateOrganizationResponse, error)
}

type UserOrganizationMembershipUsecase interface {
	DeleteMembership(executorID, deleteUserID, organizationID string) (*response.DeleteMembershipResponse, error)
}

type UserOrganizationInvitationUsecase interface {
	CreateInvitation(orgID, userID, invitedBy string) (*response.CreateInvitationResponse, error)
}
