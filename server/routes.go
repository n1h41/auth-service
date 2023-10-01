package server

import (
	"n1h41/auth-service/config"
	"n1h41/auth-service/controllers"
	"n1h41/auth-service/helpers"
	"n1h41/auth-service/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(config *config.Config) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// INFO: Initialize database
	db, err := helpers.InitDB(config)

	if err != nil {
		panic(err)
	}

  // INFO: Status endpoint
	health := controllers.NewHealthController()
	router.GET("/health", health.GetHealthStatus)

  // INFO: Initialize services and controllers
	// ****
	authService := services.NewAuthService(db)
	controllers.NewAuthController(router, authService)
	// ****

	return router
}
