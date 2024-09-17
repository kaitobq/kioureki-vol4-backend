package controller

import (
	"kioureki-vol4-backend/internal/domain/service"
	"kioureki-vol4-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(
	r *gin.Engine,
	userCtrl *UserController,
	organizationCtrl *OrganizationController,
	membershipCtrl *UserOrganizationMembershipController,
	invitationCtrl *UserOrganizationInvitationController,
	tokenService service.TokenService,
) {
	v1 := r.Group("api/v1")

	auth := v1.Group("auth")
	{
		auth.POST("/signup", userCtrl.SignUp)
		auth.POST("/signin", userCtrl.SignIn)
	}

	org := v1.Group("organization")
	org.Use(middleware.AuthMiddleware(tokenService))
	{
		org.POST("", organizationCtrl.CreateOrganization)
	}

	membership := v1.Group("membership")
	membership.Use(middleware.AuthMiddleware(tokenService))
	{
		membership.POST("", membershipCtrl.CreateMembership)
	}

	invitation := v1.Group("invitation")
	invitation.Use(middleware.AuthMiddleware(tokenService))
	{
		invitation.POST("", invitationCtrl.CreateInvitation)
	}
}
