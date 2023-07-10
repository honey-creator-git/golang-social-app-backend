package controllers

import (
	"context"
	"fmt"
	"net/http"
	"serendipity_backend/SerendipityRequest"
	"serendipity_backend/SerendipityResponse"
	"serendipity_backend/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddNewFoundation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "User is not lgged in."})
			return
		}
		role, exists := ctx.Get("role")
		user_role, _ := strconv.Atoi(fmt.Sprintf("%v", role))
		fmt.Printf("User Role for Update => %v", user_role)

		if user_role == 3 {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_PERMISSION_ALLOWED, "status": "Permission is not arranged."})
			return
		}
		user_email := fmt.Sprintf("User Email %v", email)
		fmt.Printf("User Email => %v", user_email)
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var newFoundation SerendipityRequest.SubmitFoundation
		defer cancel()

		if err := ctx.ShouldBindJSON(&newFoundation); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&newFoundation); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validating of request body."})
			return
		}

		createFoundation := models.Foundation{
			ID:   primitive.NewObjectID(),
			Name: newFoundation.Name,
		}

		curFoundation, err := createFoundation.SaveFoundation(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred for add a new foundation."})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "new_foundation": curFoundation})
	}
}

func GetAllFoundations() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "User is not lgged in."})
			return
		}
		role, exists := ctx.Get("role")
		user_role, _ := strconv.Atoi(fmt.Sprintf("%v", role))
		fmt.Printf("User Role for Update => %v", user_role)

		if user_role == 3 {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_PERMISSION_ALLOWED, "status": "Permission is not arranged."})
			return
		}
		user_email := fmt.Sprintf("User Email %v", email)
		fmt.Printf("User Email => %v", user_email)
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		results, err := models.GetAllFoundations(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": results})
	}
}
