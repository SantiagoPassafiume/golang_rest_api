package routes

import (
	"github.com/SantiagoPassafiume/golang_rest_api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Coult not save user."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created correctly."})
}
