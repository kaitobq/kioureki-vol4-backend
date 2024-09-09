package request

import "github.com/gin-gonic/gin"

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewSignUpRequest(c *gin.Context) (*SignUpRequest, error) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}
