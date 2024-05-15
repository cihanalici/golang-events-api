package routes

import (
	"net/http"

	"github.com/cihanalici/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
