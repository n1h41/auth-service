package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatController interface {
	HealthCheck(c *gin.Context)
}

type chatController struct{}

func NewChatController(router *gin.Engine) {
	controller := &chatController{}
	api := router.Group("chat")
	api.GET("/healthcheck", controller.HealthCheck)
}

func (controller *chatController) HealthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": true, "message": "Chat Endpoint live"})
}
