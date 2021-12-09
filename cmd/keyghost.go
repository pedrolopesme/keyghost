package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/pedrolopesme/keyghost/internal/core/domain"
	"github.com/pedrolopesme/keyghost/internal/handlers"
	"go.uber.org/zap"
)

type Keyghost struct {
	Ctx domain.AppContext
}

func NewKeyghost(logger *zap.Logger) *Keyghost {
	appContext := domain.NewAppContext(*logger)

	return &Keyghost{
		Ctx: *appContext,
	}
}

func (k *Keyghost) Run() error {
	logger := k.Ctx.Logger()
	logger.Info("Starting up Keyghost")

	// Configure application Routes
	e := echo.New()
	openIdRoutes := handlers.NewOpenIdGhostHandlers(k.Ctx, e)
	openIdRoutes.SetupRoutes()

	err := e.Start(":8080")
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}
