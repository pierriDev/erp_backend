package main

import (
	"github.com/pierriDev/erp_backend.git/config"
	"github.com/pierriDev/erp_backend.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//INITIALIZE CONFIGS
	err := config.Init()
	if err != nil {
		logger.ErrorF("Config initialization error: %v", err)
		return
	}

	router.Initialize()
}
