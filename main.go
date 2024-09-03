package main

import (
	"fmt"
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	server.Run(":8080") //localHost:8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvent()
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data", "newMessage": err})
		fmt.Println(err)
		return
	}

	event.ID = 1
	event.UserId = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "data sussefully created", "event": event})
}
