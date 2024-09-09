//go:build wireinject

package app

import (
	"go-template/internal/app/config"
	"go-template/internal/app/container"
	"go-template/internal/controller"
	"go-template/internal/infra/db"
	"go-template/internal/usecase"
	"go-template/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func New() (*container.App, error) {
	wire.Build(
		provideGinEngine,
		config.New,
		config.NewDBConfig,
		container.NewApp,
		database.New,
		container.NewCtrl,

		// user
		controller.NewUserController,
		usecase.NewUserUsecase,
		db.NewUserRepository,
	)

	return nil, nil
}

func provideGinEngine() *gin.Engine {
	return gin.Default()
}
