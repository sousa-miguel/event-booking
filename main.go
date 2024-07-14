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
	err := server.Run(":8080")
	if err != nil {
		return
	}
}
