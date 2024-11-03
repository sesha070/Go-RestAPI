package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"restApi/models"

	"restApi/db"
)

func main() {
	db.InitDB()
	server := gin.Default() //sets up the basic server
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080") // localhost on port 8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could fetch data from db"})
		return
	}
	context.JSON(http.StatusOK, events)

}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data , please provie all the values"})

	}

	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		fmt.Print(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not add data to db", "err": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event added", "event": event})

}
