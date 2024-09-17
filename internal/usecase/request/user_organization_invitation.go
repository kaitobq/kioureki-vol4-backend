package request

import "github.com/gin-gonic/gin"

type CreateInvitationRequest struct {
	OrganizationID string    `json:"organization_id"`
	Email		   string    `json:"email"`//招待したいユーザのemailで招待できるように変更
}

func NewCreateInvitationRequest(c *gin.Context) (*CreateInvitationRequest, error) {
	var req CreateInvitationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
