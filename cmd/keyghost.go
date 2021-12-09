package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/pedrolopesme/keyghost/internal/core/domain"
	"github.com/pedrolopesme/keyghost/internal/handlers"
	"go.uber.org/zap"
)

type Keyghost struct {
	Ctx domain.AppContext
}

func NewKeyghost(logger *zap.Logger, serverProfilePath string) *Keyghost {
	serverProfile := loadServerProfile(serverProfilePath)
	appContext := domain.NewAppContext(*logger, *serverProfile)

	return &Keyghost{
		Ctx: *appContext,
	}
}

func loadServerProfile(serverProfilePath string) *domain.ServerProfile {
	fileBytes, err := ioutil.ReadFile(serverProfilePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var serverProfile domain.ServerProfile
	err = json.Unmarshal(fileBytes, &serverProfile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return &serverProfile
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
