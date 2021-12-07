package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OpenIdGhostHandlers struct {
	ctx context.Context
}

func NewOpenIdGhostHandlers(ctx context.Context) *OpenIdGhostHandlers {
	return &OpenIdGhostHandlers{
		ctx: ctx,
	}
}

func (openid OpenIdGhostHandlers) SetupRoutes(e echo.Echo) {
	e.GET("/auth/realms/:realm/.well-known/openid-configuration", openid.wellknown)
}

func (openid OpenIdGhostHandlers) wellknown(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
