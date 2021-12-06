package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pedrolopesme/keyghost/internal/core/domain"
	"go.uber.org/zap"
)

type NewOpenIdGhostService struct {
	logger  zap.Logger
	profile domain.ServerProfile
}

func NewNewOpenIdGhostService(ctx context.Context, profile domain.ServerProfile) *NewOpenIdGhostService {
	return &NewOpenIdGhostService{
		logger:  ctx.Value("logger").(zap.Logger),
		profile: profile,
	}
}

func (o NewOpenIdGhostService) Wellknown() string {
	wellknown, err := json.Marshal(o.profile.Wellknown)
	if err != nil {
		o.logger.Error(err.Error())
		return ""
	}

	response := fmt.Sprintf(string(wellknown), o.profile.Realm)

	return response
}
