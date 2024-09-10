package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")          // group the routes together
	authenticated.Use(middlewares.Authenticate) // use middleware no all the routes in the group
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvents)
	authenticated.POST("/events/:id/register", registerForEvnet)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// server.POST()
	// server.PUT()
	// server.DELETE()

	server.POST("/signup", signup)
	server.POST("/login", login)

}
