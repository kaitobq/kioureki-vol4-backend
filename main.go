package main

import (
	"backend/controllers"
	"backend/middlewares"
	"backend/models"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	if models.DB == nil {
		log.Fatal("Database connection failed1")
	}

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(config))

	public := router.Group("/api")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}

	protected := router.Group("/api/protected")
	protected.Use(middlewares.JwtAuthMiddleware())

	{
		user := protected.Group("/user")
		{
			user.GET("", controllers.CurrentUser)
			user.PATCH("", controllers.UpdateUsername)
		}


		organization := protected.Group("/organization")
		{
			organization.GET("", controllers.GetOrganizations)
			organization.OPTIONS("", controllers.GetOrganizations)
			organization.POST("", controllers.CreateOrganization)
			organization.DELETE("/:id", controllers.DeleteOrganization)
			organization.GET("/:id", controllers.GetMembers)
		}

		invite := protected.Group("/organization/invite")
		{
			invite.GET("", controllers.GetInvitations)
			invite.POST("", controllers.CreateInvitation)
			invite.POST("/:id", controllers.AcceptInvitation)
			invite.DELETE("/:id", controllers.DeleteInvitation)
		}

		player := protected.Group("/organization/:id/player")
		{
			player.POST("", controllers.CreatePlayer)
			player.GET("", controllers.GetPlayers)
			player.DELETE("/:player_id", controllers.DeletePlayer)
			player.POST("/:player_id/category", controllers.AddCategory)
			player.GET("/category", controllers.GetPlayerCategories)
			player.GET("/:player_id/detail", controllers.GetPlayerDetail)
		}

		category := protected.Group("/organization/:id/category")
		{
			category.POST("", controllers.CreateCategory)
			category.GET("", controllers.GetCategories)
		}

		medicalrecord := protected.Group("/organization/:id/medicalrecord")
		{
			medicalrecord.POST("", controllers.CreateMedicalRecord)
			medicalrecord.GET("", controllers.GetMedicalRecords)
			medicalrecord.PATCH("/:record_id", controllers.UpdateMedicalRecord)
			medicalrecord.GET("/day", controllers.GetInjuredToday)
			medicalrecord.GET("/week", controllers.GetInjuredThisWeek)
			medicalrecord.GET("/month", controllers.GetInjuredThisMonth)
			medicalrecord.GET("/status/:status", controllers.GetRecordByStatus)
			medicalrecord.GET("/recover/week", controllers.GetRecoverThisWeek)
			medicalrecord.GET("/filter/category/:category_id/:option_id", controllers.GetFilteredByCategory)
		}
	}

	router.Run(":8080")
}
