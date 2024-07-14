package main

import (
	"example.com/rest-api-gin/db"
	"example.com/rest-api-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
