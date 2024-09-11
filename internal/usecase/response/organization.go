package response

import "time"

type OrganizationResponse struct {
	Name      string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateOrganizationResponse struct {
	Message   string `json:"message"`
	Organization OrganizationResponse `json:"organization"`
}

func NewCreateOrganizationResponse(name string, createdAt time.Time, updatedAt time.Time) (*CreateOrganizationResponse, error) {
	return &CreateOrganizationResponse{
		Message: "Organization created successfully",
		Organization: OrganizationResponse{
			Name:      name,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}, nil
}
