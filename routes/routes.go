package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // GET request
	server.GET("/events/:id",getEvent) // GET a single event by ID
	server.POST("/events", createEvent) // POST request
}