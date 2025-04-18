package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) // GET request
	server.GET("/events/:id",getEvent) // GET a single event by ID

	// middleware
	authenticated := server.Group("/")
	authenticated.Use()
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signUp) // POST request for creating new users
	server.POST("/login", login) // POST login
}