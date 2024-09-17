package response

type CreateMembershipResponse struct {
	Message string `json:"message"`
}

func NewCreateMembershipResponse() (*CreateMembershipResponse, error) {
	return &CreateMembershipResponse{
		Message: "Membership created successfully",
	}, nil
}

type DeleteMembershipResponse struct {
	Message string `json:"message"`
}

func NewDeleteMembershipResponse() (*DeleteMembershipResponse, error) {
	return &DeleteMembershipResponse{
		Message: "Membership deleted successfully",
	}, nil
}
