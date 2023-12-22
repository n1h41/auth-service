package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"n1h41/auth-service/cmd/api/controllers"
	"n1h41/auth-service/config"
	authController "n1h41/auth-service/features/auth/controllers"
	"n1h41/auth-service/features/auth/services"
	"n1h41/auth-service/internal/db"
)

func SetupRouter(config *config.Config) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// INFO: Initialize database
	db, err := db.InitDB(config)
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}

	// INFO: Status endpoint
	health := controllers.NewHealthController()
	router.GET("/health", health.GetHealthStatus)

	// INFO: Initialize services and controllers
	// ****
	authService := services.NewAuthService(db)
	authController.NewAuthController(router, authService)
	// ****

	return router
}
