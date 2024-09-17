package usecase

import (
	"fmt"
	"kioureki-vol4-backend/internal/domain/entity"
	"kioureki-vol4-backend/internal/domain/repository"
	"kioureki-vol4-backend/internal/domain/service"
	"kioureki-vol4-backend/internal/usecase/response"
)

type userOrganizationInvitationUsecase struct {
	repo repository.UserOrganizationInvitationRepository
	ulidService service.ULIDService
	userRepo repository.UserRepository
}

func NewUserOrganizationInvitationUsecase(repo repository.UserOrganizationInvitationRepository, ulidService service.ULIDService, userRepo repository.UserRepository) UserOrganizationInvitationUsecase {
	return &userOrganizationInvitationUsecase{
		repo: repo,
		ulidService: ulidService,
		userRepo: userRepo,
	}
}

func (u *userOrganizationInvitationUsecase) CreateInvitation(orgID, userEmail, invitedBy string) (*response.CreateInvitationResponse, error) {
	user, err := u.userRepo.FindByEmail(userEmail)
	if err != nil {
		fmt.Printf("find by email[%s]\n", userEmail)
		if err.Error() == fmt.Sprintf("user with email %s not found", userEmail) {
			return nil, fmt.Errorf("unable to invite user: no user with email %s found", userEmail)
		}
		return nil, err
	}
	
	id := u.ulidService.GenerateULID()
	invitation := entity.UserOrganizationInvitation{
		ID: id,
		OrganizationID: orgID,
		UserID: user.ID,
	}

	creator, err := u.userRepo.FindByID(invitedBy)
	if err != nil {
		fmt.Println("find by id")
		return nil, err
	}

	err = u.repo.CreateInvitation(invitation, creator.Name)
	if err != nil {
		fmt.Println("create invitation")
		return nil, err
	}

	return response.NewCreateInvitationResponse()
}
