package container

import (
	"errors"
	"fmt"
	"go-template/internal/app/config"
	"go-template/internal/controller"
	"go-template/pkg/database"

	"github.com/gin-gonic/gin"
)

type container struct {
	userCtrl *controller.UserController
}

func NewCtrl(
	userCtrl *controller.UserController,
) *container {
	return &container{
		userCtrl: userCtrl,
	}
}

type App struct {
	r   *gin.Engine
	cfg *config.Config
	db  *database.DB
}

func NewApp(r *gin.Engine, container *container, cfg *config.Config, db *database.DB) *App {
	controller.SetUpRoutes(r, container.userCtrl)

	return &App{
		r: r,
		cfg: cfg,
		db: db,
	}
}

func (a *App) Migrate() error {
	return a.db.Migrate()
}

func (a *App) Run() error {
	return a.r.Run(fmt.Sprintf(":%d", a.cfg.Port))
}

func (a *App) Close() error {
	return errors.Join(
		a.db.Close(),
	)
}
