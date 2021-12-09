package main

import (
	"flag"

	"github.com/pedrolopesme/keyghost/cmd"
	"go.uber.org/zap"
)

func main() {
	serverProfilePath := flag.String("profile", "", "where is the server profile json file?")
	flag.Parse()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	keyghost := cmd.NewKeyghost(logger, *serverProfilePath)
	keyghost.Run()
}
