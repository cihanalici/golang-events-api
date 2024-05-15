package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)          // Get all events
	server.GET("/events/:id", getEvent)       // Get a single event
	server.POST("/events", createEvent)       // Create a new event
	server.PUT("/events/:id", updateEvent)    // Update an event
	server.DELETE("/events/:id", deleteEvent) // Delete an event
	server.POST("/signup", signup)            // Create a new user
}
