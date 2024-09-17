package entity

type UserOrganizationInvitation struct {
	ID             string
	OrganizationID string
	UserID		   string
	InvitedBy      string
	CreatedAt      string
	UpdatedAt      string
}