package response

type CreateInvitationResponse struct {
	Message string `json:"message"`
}

func NewCreateInvitationResponse() (*CreateInvitationResponse, error) {
	return &CreateInvitationResponse{
		Message: "Invitation created successfully",
	}, nil
}
