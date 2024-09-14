package controller

import "github.com/gin-gonic/gin"

func SetUpRoutes(
	r *gin.Engine,
	userCtrl *UserController,
	organizationCtrl *OrganizationController,
	userOrganizationMembershipCtrl *UserOrganizationMembershipController,
) {
	v1 := r.Group("api/v1")

	auth := v1.Group("auth")
	{
		auth.POST("/signup", userCtrl.SignUp)
		auth.POST("/signin", userCtrl.SignIn)
	}

	org := v1.Group("organization")
	{
		org.POST("", organizationCtrl.CreateOrganization)
	}

	membership := v1.Group("membership")
	{
		membership.POST("", userOrganizationMembershipCtrl.CreateMembership)
	}
}
