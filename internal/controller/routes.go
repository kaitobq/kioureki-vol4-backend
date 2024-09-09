package controller

import "github.com/gin-gonic/gin"

func SetUpRoutes(
	r *gin.Engine,
	userCtrl *UserController,
) {
	v1 := r.Group("api/v1")

	auth := v1.Group("auth")
	{
		auth.POST("/signup", userCtrl.SignUp)
		auth.POST("/signin", userCtrl.SignIn)
	}
}
