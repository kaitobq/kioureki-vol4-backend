package request

import "github.com/gin-gonic/gin"

type CreateMembershipRequest struct {
	UserID         string `json:"user_id" binding:"required"`
	OrganizationID string `json:"organization_id" binding:"required"`
}

func NewCreateMembershipRequest(c *gin.Context) (*CreateMembershipRequest, error) {
	var req CreateMembershipRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

type DeleteMembershipRequest struct {
	UserID         string `json:"user_id" binding:"required"`
	OrganizationID string `json:"organization_id" binding:"required"`
}

func NewDeleteMembershipRequest(c *gin.Context) (*DeleteMembershipRequest, error) {
	var req DeleteMembershipRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
