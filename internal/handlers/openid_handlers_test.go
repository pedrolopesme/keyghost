package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenIdConnect(t *testing.T) {

	// mockedLogger := zaptest.NewLogger(t)
	// mockedServerProfile := domain.ServerProfile{}
	// appContextMock := domain.NewAppContext(
	// 	*mockedLogger,
	// 	mockedServerProfile,
	// )

	t.Run("Wellknown", func(t *testing.T) {

		t.Run("When X then Y", func(t *testing.T) {
			assert.Nil(t, true)
		})
	})
}
