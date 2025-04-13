package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // set up Engine (HTTP server) -> pointer *gin.Engine

	server.GET("/events", getEvents) // GET request

	server.Run(":8080") // listening for incoming requests on localhost:8080
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello"})
}