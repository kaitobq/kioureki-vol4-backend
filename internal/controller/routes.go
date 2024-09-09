package controller

import "github.com/gin-gonic/gin"

func SetUpRoutes(
	r *gin.Engine,
	userCtrl *UserController,
) {
	r.POST("/signup", userCtrl.SignUp)
}
