package main

import (
	"github.com/pedrolopesme/keyghost/cmd"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	keyghost := cmd.NewKeyghost(logger)
	keyghost.Run()
}
