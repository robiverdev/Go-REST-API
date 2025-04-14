package main

import (
	"log"
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default() // set up Engine (HTTP server) -> pointer *gin.Engine

	server.GET("/events", getEvents) // GET request
	server.GET("/events/:id",getEvent) // GET a single event
	server.POST("/events", createEvent) // POST request

	server.Run(":8080") // listening for incoming requests on localhost:8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not parse request data"})
		return
	}

	// dummy value
	event.ID = 1
	event.UserID = 1

	err =event.Save()
	if err != nil {
		log.Println("Error saving event:", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event" : event})
}