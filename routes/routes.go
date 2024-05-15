package routes

import (
	"github.com/cihanalici/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)    // Get all events
	server.GET("/events/:id", getEvent) // Get a single event

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)                   // Create a new event
	authenticated.PUT("/events/:id", updateEvent)                // Update an event
	authenticated.DELETE("/events/:id", deleteEvent)             // Delete an event
	authenticated.POST("/events/:id/register", registerForEvent) // Register for an event
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup) // Create a new user
	server.POST("/login", login)   // Login user
}
