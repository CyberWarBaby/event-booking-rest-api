package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    // GET, POST, PUT, DELETE, PATCH
	server.GET("/events/:id", getEvent) // /event/1
	server.POST("/events", createEvent)
	server.PUT("/events/:id")
}
