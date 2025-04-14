package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default() // set up Engine (HTTP server) -> pointer *gin.Engine
	routes.RegisterRoutes(server)
	
	server.Run(":8080") // listening for incoming requests on localhost:8080
}