package server

import (
	"n1h41/auth-service/config"
)

func InitServer() {
	// INFO: Load config
	config, err := config.LoadConfig("./")
	if err != nil {
		panic(err)
	}

	// INFO: Setup router
	router := SetupRouter(config)

	port := ":" + config.Port
	router.Run("0.0.0.0" + port)
}
