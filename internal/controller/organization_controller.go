package controller

import (
	"go-template/internal/usecase"
	"go-template/internal/usecase/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	uc usecase.OrganizationUsecase
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
	
	res, err := ct.uc.CreateOrganization(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}
