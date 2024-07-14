package routes

import (
	"example.com/rest-api-gin/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	// authenticated paths
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// users
	server.POST("/signup", signup)
	server.POST("/login", login)
}
