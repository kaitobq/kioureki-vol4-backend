package controller

import (
	"kioureki-vol4-backend/internal/domain/service"
	"kioureki-vol4-backend/internal/usecase"
	"kioureki-vol4-backend/internal/usecase/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	uc usecase.OrganizationUsecase
	tokenService service.TokenService
}

func NewOrganizationController(uc usecase.OrganizationUsecase, tokenService service.TokenService) *OrganizationController {
	return &OrganizationController{
		uc: uc,
		tokenService: tokenService,
	}
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
