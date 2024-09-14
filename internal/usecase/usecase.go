package usecase

import "go-template/internal/usecase/response"

type UserUsecase interface {
	CreateUser(name, email, password string) (*response.SignUpResponse, error)
	SignIn(email, password string) (*response.SignInResponse, error)
}

type OrganizationUsecase interface {
	CreateOrganization(name string, founderID uint) (*response.CreateOrganizationResponse, error)
}

type UserOrganizationMembershipUsecase interface {
	CreateMembership(userID, organizationID uint) (*response.CreateMembershipResponse, error)
}
