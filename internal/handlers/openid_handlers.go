package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pedrolopesme/keyghost/internal/core/domain"
)

type OpenIdGhostHandlers struct {
	ctx domain.AppContext
	e   *echo.Echo
}

func NewOpenIdGhostHandlers(ctx domain.AppContext, e *echo.Echo) *OpenIdGhostHandlers {
	return &OpenIdGhostHandlers{
		ctx: ctx,
		e:   e,
	}
}

func (openid OpenIdGhostHandlers) SetupRoutes() {
	openid.e.GET("/auth/realms/:realm/.well-known/openid-configuration", openid.wellknown)
}

func (openid OpenIdGhostHandlers) wellknown(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
