package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"iqge.org/gym/controllers"
	"iqge.org/gym/utils"
)

var db = make(map[string]string)

func initMongoClient() {
	_, err := utils.Connect()
	if err != nil {
		log.Fatal(err)
	}
}
func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	initMongoClient()

	// Add unhandled rejection handler
	r.Use(func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Handle the rejection here
				utils.LogError("RejectionMiddleware", errors.New(fmt.Sprintf("%v", r)))
				c.JSON(500, gin.H{"error": "Something went wrong"})
				c.Abort()
			}
		}()
		c.Next()
	})

	healthCheckController := controllers.HealthCheckController{}
	gymController := controllers.GymController{}
	// Ping test
	r.GET("/ping", healthCheckController.Index)
	r.GET("/pong", gymController.CreateGym)
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080

	r.Run(":8080")
}
