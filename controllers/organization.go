package controllers

import (
	"backend/models"
	"backend/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateOrganizationInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateOrganization(c *gin.Context) {
	var input CreateOrganizationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organization := models.Organization{Name: input.Name}
	if err := models.DB.Create(&organization).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := token.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	membership := models.Membership{OrganizationID: organization.ID, UserID: userId}
	models.DB.Create(&membership)

	c.JSON(http.StatusOK, gin.H{"data": organization})
}

type AddUserToOrganizationInput struct {
	OrganizationID uint `json:"organization_id" binding:"required"`
	UserID		 uint `json:"user_id" binding:"required"`
}

func AddUserToOrganization(c *gin.Context) {
	var input AddUserToOrganizationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	membership := models.Membership{OrganizationID: input.OrganizationID, UserID: input.UserID}
	if err := models.DB.Create(&membership).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": membership})
}

func GetOrganizations(c *gin.Context) {
	userId, err := token.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var memberships []models.Membership
	if err = models.DB.Where("user_id = ?", userId).Find(&memberships).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var organizations []models.Organization
	for _, membership := range memberships {
		var organization models.Organization
		models.DB.First(&organization, membership.OrganizationID)
		organizations = append(organizations, organization)
	}

	c.JSON(http.StatusOK, gin.H{"organizations": organizations})
}

func GetMembers(c *gin.Context){
	organizationId := c.Param("id")
	var memberships []models.Membership
	if err := models.DB.Where("organization_id = ?", organizationId).Find(&memberships).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var response []map[string]interface{}
	for _, membership := range memberships {
		var user models.User
		models.DB.First(&user, membership.UserID)
		response = append(response, map[string]interface{}{
			"id": membership.ID,
			"user": user.Username,
			"joined_at": membership.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"members": response})
}

func DeleteOrganization(c *gin.Context) {
	organizationId := c.Param("id")
	if err := models.DB.Transaction(func(tx *gorm.DB) error {
		// 組織の削除
		if err := tx.Where("id = ?", organizationId).Delete(&models.Organization{}).Error; err != nil {
			return err
		}

		// 関連するメンバーシップの削除
		if err := tx.Where("organization_id = ?", organizationId).Delete(&models.Membership{}).Error; err != nil {
			return err
		}

		// 関連する招待の削除
		if err := tx.Where("organization_id = ?", organizationId).Delete(&models.Invitation{}).Error; err != nil {
			return err
		}

		// 関連するプレイヤーの削除
		if err := tx.Where("organization_id = ?", organizationId).Delete(&models.Player{}).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Organization deleted"})
}
