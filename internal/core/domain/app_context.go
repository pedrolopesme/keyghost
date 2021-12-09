package domain

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

const (
	CTX_LOGGER         = "LOGGER"
	CTX_SERVER_PROFILE = "SERVER_PROFILE"
)

type AppContext struct {
	ctx           context.Context
	serverProfile ServerProfile
}

func NewAppContext(logger zap.Logger, serverProfile ServerProfile) *AppContext {
	ctx := context.Background()
	ctx = context.WithValue(ctx, CTX_LOGGER, logger)
	ctx = context.WithValue(ctx, CTX_SERVER_PROFILE, serverProfile)

	return &AppContext{
		ctx: ctx,
	}
}

func (app AppContext) Logger() zap.Logger {
	return app.ctx.Value(CTX_LOGGER).(zap.Logger)
}

func (app AppContext) ServerProfile() ServerProfile {
	return app.ctx.Value(CTX_SERVER_PROFILE).(ServerProfile)
}
