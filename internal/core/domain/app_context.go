package domain

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

const (
	CTX_LOGGER = "LOGGER"
)

type AppContext struct {
	ctx context.Context
}

func NewAppContext(logger zap.Logger) *AppContext {
	ctx := context.Background()
	ctx = context.WithValue(ctx, CTX_LOGGER, logger)

	return &AppContext{
		ctx: ctx,
	}
}

func (app AppContext) Logger() zap.Logger {
	return app.ctx.Value(CTX_LOGGER).(zap.Logger)
}
