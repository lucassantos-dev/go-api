package main

import (
	"github.com/lucassantos-dev/go-api/config"
	"github.com/lucassantos-dev/go-api/router"
)

var (
	logger *config.Logger
)

func main() {
	err := config.Init()

	logger = config.GetLogger("main")
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}
	router.Initializer()
}
