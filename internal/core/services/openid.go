package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pedrolopesme/keyghost/internal/core/domain"
	"go.uber.org/zap"
)

type OpenIdGhost struct {
	logger  zap.Logger
	profile domain.ServerProfile
}

func NewOpenIdGhost(ctx context.Context, profile domain.ServerProfile) *OpenIdGhost {
	return &OpenIdGhost{
		logger:  ctx.Value("logger").(zap.Logger),
		profile: profile,
	}
}

func (o OpenIdGhost) Wellknown() string {
	wellknown, err := json.Marshal(o.profile.Wellknown)
	if err != nil {
		o.logger.Error(err.Error())
		return ""
	}

	response := fmt.Sprintf(string(wellknown), o.profile.Realm)

	return response
}
