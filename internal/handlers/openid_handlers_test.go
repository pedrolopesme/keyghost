package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/pedrolopesme/keyghost/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestOpenIdConnect(t *testing.T) {

	var (
		serverProfileMock  *domain.ServerProfile
		appContextMock     *domain.AppContext
		echoMock           *echo.Echo
		openIdHandlersMock OpenIdGhostHandlers
	)

	setupTest := func(t *testing.T) func(t *testing.T) {
		appContextMock = domain.NewAppContext(*zaptest.NewLogger(t), *serverProfileMock)
		echoMock = echo.New()

		openIdHandlersMock = *NewOpenIdGhostHandlers(appContextMock, echoMock)
		openIdHandlersMock.SetupRoutes()

		return func(t *testing.T) {
			appContextMock = nil
			serverProfileMock = nil
			echoMock = nil
		}
	}
	// mockedLogger := zaptest.NewLogger(t)
	// mockedServerProfile := domain.ServerProfile{}
	// appContextMock := domain.NewAppContext(
	// 	*mockedLogger,
	// 	mockedServerProfile,
	// )

	t.Run("Wellknown", func(t *testing.T) {

		t.Run("WhenCallWitARealmThatNotExistsThenReturnNotFound", func(t *testing.T) {
			serverProfileMock = &domain.ServerProfile{}
			defer setupTest(t)(t)

			req := httptest.NewRequest(http.MethodGet, "/auth/realms/realm-not-found/.well-known/openid-configuration", nil)
			rec := httptest.NewRecorder()
			reqContext := echoMock.NewContext(req, rec)

			if assert.NoError(t, openIdHandlersMock.wellknown(reqContext)) {
				assert.Equal(t, http.StatusNotFound, rec.Code)
			}
		})

	})
}
