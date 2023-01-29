package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckController struct{}

// Index handles the GET / route and displays a list of users
func (ctrl HealthCheckController) Index(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"hello": "There"})
}
