package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController interface {
	GetHealthStatus(c *gin.Context) // GET /health
}

type healthController struct{}

func (controller healthController) GetHealthStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}

func NewHealthController() HealthController {
	return &healthController{}
}
