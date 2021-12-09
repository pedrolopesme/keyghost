package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pedrolopesme/keyghost/internal"
	"github.com/pedrolopesme/keyghost/internal/core/domain"
)

type OpenIdGhostHandlers struct {
	ctx           domain.AppContext
	e             *echo.Echo
	serverProfile domain.ServerProfile
}

func NewOpenIdGhostHandlers(ctx domain.AppContext, e *echo.Echo, serverProfile domain.ServerProfile) *OpenIdGhostHandlers {
	return &OpenIdGhostHandlers{
		ctx:           ctx,
		e:             e,
		serverProfile: serverProfile,
	}
}

func (openid OpenIdGhostHandlers) SetupRoutes() {
	openid.e.GET("/auth/realms/:realm/.well-known/openid-configuration", openid.wellknown)
}

func (openid OpenIdGhostHandlers) wellknown(c echo.Context) error {
	realmParam := c.Param("realm")
	realm := openid.serverProfile.Realm(realmParam)
	if realm == nil {
		return internal.ErrRealmNotFound
	}

	payloadBytes, err := json.Marshal(realm.Wellknown)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, string(payloadBytes))
}
