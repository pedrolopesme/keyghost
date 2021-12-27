package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pedrolopesme/keyghost/internal"
	"github.com/pedrolopesme/keyghost/internal/core/domain"
	"go.uber.org/zap"
)

const (
	GRANT_TYPE_AUTHORIZATION = "authorization_code"
)

type OpenIdGhostHandlers struct {
	ctx *domain.AppContext
	e   *echo.Echo
}

func NewOpenIdGhostHandlers(ctx *domain.AppContext, e *echo.Echo) *OpenIdGhostHandlers {
	return &OpenIdGhostHandlers{
		ctx: ctx,
		e:   e,
	}
}

func (openid OpenIdGhostHandlers) SetupRoutes() {
	openid.e.GET("/auth/realms/:realm/.well-known/openid-configuration", openid.wellknown)
	openid.e.GET("/auth/realms/:realm/protocol/openid-connect/auth", openid.authorization)
	openid.e.POST("/auth/realms/:realm/protocol/openid-connect/token", openid.token)
	openid.e.GET("/redirect/callback", openid.authorizationCallback)
}

func (openid OpenIdGhostHandlers) wellknown(c echo.Context) error {
	logger := openid.ctx.Logger()
	realmParam := c.Param("realm")
	realm := openid.ctx.ServerProfile().Realm(realmParam)
	if realm == nil {
		logger.Info("Realm not found", zap.Error(internal.ErrRealmNotFound))
		return c.NoContent(http.StatusNotFound)
	}

	payloadBytes, err := json.Marshal(realm.Wellknown)
	if err != nil {
		return err
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return c.String(http.StatusOK, string(payloadBytes))
}

func (openid OpenIdGhostHandlers) authorization(c echo.Context) error {
	realmParam := c.Param("realm")
	realm := openid.ctx.ServerProfile().Realm(realmParam)
	if realm == nil {
		return internal.ErrRealmNotFound
	}

	// preparing return URL parameters (code and state)
	redirectUrl := realm.Authorization.RedirectTo
	if strings.Contains(redirectUrl, "?") {
		redirectUrl += "&code=%s"
	} else {
		redirectUrl += "?code=%s"
	}
	redirectUrl = fmt.Sprintf(redirectUrl, realm.Authorization.Code)

	if c.QueryParam("state") != "" {
		redirectUrl += "&state=" + c.QueryParam("state")
	}

	return c.Redirect(http.StatusFound, redirectUrl)
}

func (openid OpenIdGhostHandlers) authorizationCallback(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func (openId OpenIdGhostHandlers) token(c echo.Context) error {
	if c.FormValue("grant_type") == GRANT_TYPE_AUTHORIZATION {
		return openId.authorizationCode(c)
	}

	return c.NoContent(http.StatusBadRequest)
}

func (openId OpenIdGhostHandlers) authorizationCode(c echo.Context) error {
	realmParam := c.Param("realm")
	realm := openId.ctx.ServerProfile().Realm(realmParam)
	if realm == nil {
		return internal.ErrRealmNotFound
	}

	if c.FormValue("code") == realm.Token.AuthorizationCode.Code {
		return c.JSON(http.StatusOK, realm.Token.AuthorizationCode.Response)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return c.NoContent(http.StatusNotFound)
}
