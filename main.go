package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello!")
	server := gin.Default()

	server.GET("/events", getEvents) // GET, POST, PUT, DELETE, PATCH
	server.Run(":8080")              //localhost:8080
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"messages": "hello!"})
}
