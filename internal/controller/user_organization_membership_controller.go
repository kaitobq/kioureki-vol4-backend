package controller

import (
	"kioureki-vol4-backend/internal/usecase"
	"kioureki-vol4-backend/internal/usecase/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserOrganizationMembershipController struct {
	uc usecase.UserOrganizationMembershipUsecase
}

func NewUserOrganizationMembershipController(uc usecase.UserOrganizationMembershipUsecase) *UserOrganizationMembershipController {
	return &UserOrganizationMembershipController{
		uc: uc,
	}
}

func (ct *UserOrganizationMembershipController) CreateMembership(c *gin.Context) {
	req, err := request.NewCreateMembershipRequest(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	res, err := ct.uc.CreateMembership(req.UserID, req.OrganizationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
