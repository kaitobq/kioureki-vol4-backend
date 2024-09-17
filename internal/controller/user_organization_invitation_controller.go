package controller

import (
	"kioureki-vol4-backend/internal/domain/service"
	"kioureki-vol4-backend/internal/usecase"
	"kioureki-vol4-backend/internal/usecase/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserOrganizationInvitationController struct {
	uc usecase.UserOrganizationInvitationUsecase
	tokenService service.TokenService
}

func NewUserOrganizationInvitationController(uc usecase.UserOrganizationInvitationUsecase, tokenService service.TokenService) *UserOrganizationInvitationController {
	return &UserOrganizationInvitationController{
		uc: uc,
		tokenService: tokenService,
	}
}

func (ct *UserOrganizationInvitationController) CreateInvitation(c *gin.Context) {
	req, err := request.NewCreateInvitationRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	id, err := ct.tokenService.ExtractIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, err := ct.uc.CreateInvitation(req.OrganizationID, req.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
