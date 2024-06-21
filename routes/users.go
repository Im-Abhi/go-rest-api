package routes

import (
	"net/http"

	"github.com/Im-Abhi/leaning-go/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't parse request data"})
		return
	}

	user.ID = 1

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created!"})
}
