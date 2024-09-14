// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/gin-gonic/gin"
	"kioureki-vol4-backend/internal/app/config"
	"kioureki-vol4-backend/internal/app/container"
	"kioureki-vol4-backend/internal/controller"
	"kioureki-vol4-backend/internal/domain/service"
	"kioureki-vol4-backend/internal/infra/db"
	"kioureki-vol4-backend/internal/usecase"
	"kioureki-vol4-backend/pkg/database"
)

// Injectors from wire.go:

func New() (*container.App, error) {
	engine := provideGinEngine()
	dbConfig := config.NewDBConfig()
	databaseDB, err := database.New(dbConfig)
	if err != nil {
		return nil, err
	}
	userRepository := db.NewUserRepository(databaseDB)
	tokenService := service.NewTokenService()
	ulidService := service.NewULIDService()
	userUsecase := usecase.NewUserUsecase(userRepository, tokenService, ulidService)
	userController := controller.NewUserController(userUsecase)
	organizationRepository := db.NewOrganizationRepository(databaseDB)
	userOrganizationMembershipRepository := db.NewUserOrganizationMembershipRepository(databaseDB)
	organizationUsecase := usecase.NewOrganizationUsecase(organizationRepository, userOrganizationMembershipRepository, ulidService)
	organizationController := controller.NewOrganizationController(organizationUsecase, tokenService)
	userOrganizationMembershipUsecase := usecase.NewUserOrganizationMembershipUsecase(userOrganizationMembershipRepository)
	userOrganizationMembershipController := controller.NewUserOrganizationMembershipController(userOrganizationMembershipUsecase)
	containerContainer := container.NewCtrl(userController, organizationController, userOrganizationMembershipController)
	configConfig := config.New()
	app := container.NewApp(engine, containerContainer, configConfig, databaseDB)
	return app, nil
}

// wire.go:

func provideGinEngine() *gin.Engine {
	return gin.Default()
}
