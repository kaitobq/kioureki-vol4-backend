package controllers

import (
	"backend/models"
	"backend/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateInvitationInput struct {
	OrganizationID uint   `json:"organization_id" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func CreateInvitation(c *gin.Context) {
	var input CreateInvitationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 招待するユーザのIDを取得
	userId, err := token.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// 招待するユーザが組織に所属しているか確認
	var membership models.Membership
	if err := models.DB.Where("user_id = ? AND organization_id = ?", userId, input.OrganizationID).First(&membership).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not a member of this organization"})
		return
	}


	// 招待されるユーザをEmailで特定
	var invitedUser models.User
	if err := models.DB.Where("email = ?", input.Email).First(&invitedUser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	//既に招待が存在するか確認
	var existingInvitation models.Invitation
	if err := models.DB.Where("organization_id = ? AND user_id = ?", input.OrganizationID, invitedUser.ID).First(&existingInvitation).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invitation already exists"})
		return
	}

	// 招待されるユーザがすでに組織に所属していないか確認
	var existingMembership models.Membership
	if err := models.DB.Where("user_id = ? AND organization_id = ?", invitedUser.ID, input.OrganizationID).First(&existingMembership).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is already a member of this organization"})
		return
	}

	// 招待を作成
	invitation := models.Invitation{OrganizationID: input.OrganizationID, UserID: invitedUser.ID}
	if err := models.DB.Create(&invitation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": invitation})
}


func GetInvitations(c *gin.Context) {
	userId, err := token.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	var invitations []models.Invitation
	if err := models.DB.Where("user_id = ?", userId).Find(&invitations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []map[string]interface{}
	for _, invitation := range invitations {
		var organization models.Organization
		if err := models.DB.First(&organization, invitation.OrganizationID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		response = append(response, map[string]interface{}{
			"id":              invitation.ID,
			"organization_id": invitation.OrganizationID,
			"organization":    organization.Name,
			"invited_at":      invitation.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"invites": response})
}

// type AcceptInvitationInput struct {
// 	OrganizationID uint `json:"organization_id" binding:"required"`
// }

func AcceptInvitation(c *gin.Context) {
	userId, err := token.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	organizationId := c.Param("id")
	var invitation models.Invitation
	if err := models.DB.Where("organization_id = ? AND user_id = ?", organizationId, userId).First(&invitation).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invitation not found"})
		return
	}

	if invitation.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to accept this invitation"})
	}

	// ユーザを組織に追加
	membership := models.Membership{OrganizationID: invitation.OrganizationID, UserID: userId}
	if err := models.DB.Create(&membership).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 招待を削除
	if err := models.DB.Delete(&invitation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Invitation accepted"})
}

func DeleteInvitation(c *gin.Context) {
	userId, err := token.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	invitationId := c.Param("id")

	var invitation models.Invitation
	if err := models.DB.First(&invitation, invitationId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invitation not found"})
		return
	}

	if invitation.UserID != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this invitation"})
		return
	}

	if err := models.DB.Delete(&invitation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Invitation deleted"})
}
