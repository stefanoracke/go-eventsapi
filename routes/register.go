package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func resgisterForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "registered"})
}


func cancelRegistartion(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	var event models.Event
	event.Id = eventId
	err = event.CancelRegistartion(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel register for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})
}
