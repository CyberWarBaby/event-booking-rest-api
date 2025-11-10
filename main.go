package main

import (
	"fmt"
	"net/http"

	"example.com/eventbookingrestapi/db"
	"example.com/eventbookingrestapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello!")
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, PUT, DELETE, PATCH
	server.POST("/events", createEvent)

	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"})
	}
	// context.JSON(http.StatusOK, gin.H{"messages": "hello!"})
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse req data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again later"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "successfully created event", "event": event})
}
