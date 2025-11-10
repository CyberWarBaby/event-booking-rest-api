package main

import (
	"fmt"

	"example.com/eventbookingrestapi/db"
	"example.com/eventbookingrestapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello!")
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
