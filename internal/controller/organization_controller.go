package controller

import (
	"go-template/internal/domain/service"
	"go-template/internal/usecase"
	"go-template/internal/usecase/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	uc usecase.OrganizationUsecase
	tokenService service.TokenService
}

func NewOrganizationController() *OrganizationController {
	return &OrganizationController{}
}

// TODO: setting middleware
func (ct *OrganizationController) CreateOrganization(c *gin.Context) {
	req, err := request.NewCreateOrganizationRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := ct.tokenService.ExtractIDFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	
	res, err := ct.uc.CreateOrganization(req.Name, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}
