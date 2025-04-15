package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // GET request
	server.GET("/events/:id",getEvent) // GET a single event by ID
	server.POST("/events", createEvent) // POST request
	server.PUT("/events/:id", updateEvent) // PUT request
	server.DELETE("/events/:id", deleteEvent) // DELETE request
	server.POST("signup", signUp) // POST request for creating new users
}