package main

import (
	"github.com/LucasBiazon/Gopportunitie.git/config"
	"github.com/LucasBiazon/Gopportunitie.git/router"
)

var (
	logger config.Logger
)

func main() {
	logger := config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}
	router.Initialize()
}
