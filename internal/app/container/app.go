package container

import (
	"errors"
	"fmt"
	"kioureki-vol4-backend/internal/app/config"
	"kioureki-vol4-backend/internal/controller"
	"kioureki-vol4-backend/internal/domain/service"
	"kioureki-vol4-backend/pkg/database"

	"github.com/gin-gonic/gin"
)

type container struct {
	userCtrl         *controller.UserController
	organizationCtrl *controller.OrganizationController
	membershipCtrl *controller.UserOrganizationMembershipController
	invitationCtrl *controller.UserOrganizationInvitationController
	tokenService service.TokenService
}

func NewCtrl(
	userCtrl         *controller.UserController,
	organizationCtrl *controller.OrganizationController,
	membershipCtrl *controller.UserOrganizationMembershipController,
	invitationCtrl *controller.UserOrganizationInvitationController,
	tokenService service.TokenService,
) *container {
	return &container{
		userCtrl: userCtrl,
		organizationCtrl: organizationCtrl,
		membershipCtrl: membershipCtrl,
		invitationCtrl: invitationCtrl,
		tokenService: tokenService,
	}
}

type App struct {
	r   *gin.Engine
	cfg *config.Config
	db  *database.DB
}

func NewApp(r *gin.Engine, container *container, cfg *config.Config, db *database.DB) *App {
	controller.SetUpRoutes(
		r,
		container.userCtrl,
		container.organizationCtrl,
		container.membershipCtrl,
		container.invitationCtrl,
		container.tokenService,
	)

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
