package services

import (
	"context"

	"github.com/pedrolopesme/keyghost/internal/core/domain"
)

type OpenIdGhost struct {
	ctx     context.Context
	profile domain.ServerProfile
}

func NewOpenIdGhost(ctx context.Context, profile domain.ServerProfile) *OpenIdGhost {
	return &OpenIdGhost{
		ctx:     ctx,
		profile: profile,
	}
}
