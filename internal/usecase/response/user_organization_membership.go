package response

type CreateMembershipResponse struct {
	Message string `json:"message"`
}

func NewCreateMembershipResponse() (*CreateMembershipResponse, error) {
	return &CreateMembershipResponse{
		Message: "Membership created successfully",
	}, nil
}
