package controllers

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"serendipity_backend/SerendipityRequest"
	"serendipity_backend/SerendipityResponse"
	"serendipity_backend/configs"
	"serendipity_backend/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var marketplaceCollection *mongo.Collection = configs.GetCollection(configs.DB, "marketplace")

func CreatNewMarketplace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
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
		var input SerendipityRequest.AddNewMarketplace
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error Occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error Occurred in validation of request."})
			return
		}

		newMarketplaceType := models.MarketPlace{
			ID:               primitive.NewObjectID(),
			Title:            input.Title,
			CoverLetterImage: input.CoverLetterImage,
			Type:             input.Type,
		}

		curMarketplace, err := newMarketplaceType.SaveMarketplace(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in creating a new marketplace"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": curMarketplace})
	}
}

func GetAllMarketplaces() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		defer cancel()

		results, err := models.GetAllMarketplaces(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": results})
	}
}

func CreateNewMarketplaceItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
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
		var input SerendipityRequest.AddNewMarketplaceItem
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error Occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error Occurred in validation of request."})
			return
		}

		newMarketplaceItem := models.MarketplaceItem{
			ID:              primitive.NewObjectID(),
			Title:           input.Title,
			Logo:            input.Logo,
			Description:     input.Description,
			Link:            input.Link,
			MarketplaceType: input.MarketplaceType,
		}

		curMarketplaceItem, err := newMarketplaceItem.SaveMarketplaceItem(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in creating a new marketplace"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": curMarketplaceItem})
	}
}

func GetAllMarketplaceItemsWithType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		marketplaceId := ctx.Param("marketplaceId")
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		defer cancel()

		objMarketplaceId, err := strconv.Atoi(marketplaceId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		limit := ctx.Request.URL.Query().Get("results")
		page := ctx.Request.URL.Query().Get("page")
		sortField := ctx.Request.URL.Query().Get("sortField")
		sortOrder := ctx.Request.URL.Query().Get("sortOrder")

		convertedLimit, er := strconv.Atoi(limit)
		if er != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": er.Error(), "status": "failed"})
			return
		}
		convertedPage, err := strconv.Atoi(page)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "failed"})
			return
		}
		result, errr := models.GetAllMarketplaceItemsWithType(convertedLimit, convertedPage, sortField, sortOrder, objMarketplaceId, c)

		if errr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errr.Error(), "status": "Error occurred in getting marketplace items."})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": result})
	}
}

func UpdateMarketPlace() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "User is not logged in."})
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
		marketplaceId := ctx.Param("marketplaceId")
		var marketplaceType SerendipityRequest.UpdateMarketplaceRequest
		defer cancel()

		objMarketplaceId, _ := primitive.ObjectIDFromHex(marketplaceId)

		if err := ctx.ShouldBindJSON(&marketplaceType); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&marketplaceType); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validation of request."})
			return
		}

		result, err := models.GetMarketplaceTypeWithID(objMarketplaceId, c)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Could not get marketplace type with id."})
		}
		if result.ID.Hex() == "000000000000000000000000" {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "The Marketplace is not existed"})
			return
		}

		t := reflect.TypeOf(marketplaceType)
		names := make([]string, t.NumField())
		for i := range names {
			pX := getAttr(&marketplaceType, t.Field(i).Name)
			fmtedPX := fmt.Sprintf("%v", pX)
			if (len(fmtedPX) != 0) && (len(fmtedPX) != 2) {
				names[i] = t.Field(i).Name
				updateMarketplaceType := bson.M{
					strings.ToLower(names[i]): pX.Interface(),
				}
				result, err := models.UpdateMarketplaceType(objMarketplaceId, updateMarketplaceType, c)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in updating marketplace type."})
					return
				}
				fmt.Printf("Result of Update Marketplace Type => %v", result)
			}
		}
		updatedMarketplaceType, errr := models.GetMarketplaceTypeWithID(objMarketplaceId, c)
		if errr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Could not get updated marketplace type with id."})
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedMarketplaceType})
	}
}

func UpdateMarketplacePost() gin.HandlerFunc {
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
		itemId := ctx.Param("itemId")
		var marketplaceItem SerendipityRequest.UpdateMarketplaceItemRequest
		defer cancel()

		objItemId, _ := primitive.ObjectIDFromHex(itemId)

		if err := ctx.ShouldBindJSON(&marketplaceItem); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&marketplaceItem); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validating of request."})
			return
		}

		getMarketplaceItem, err := models.GetMarketplaceItemOne(objItemId, c)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Could not get Marketplace Item."})
			return
		}

		fmt.Printf("Get Toolkit Post Result => %v", getMarketplaceItem)

		t := reflect.TypeOf(marketplaceItem)
		names := make([]string, t.NumField())
		for i := range names {
			pX := getAttr(&marketplaceItem, t.Field(i).Name)
			fmtedPX := fmt.Sprintf("%v", pX)
			if pX.Interface() != 0 && len(fmtedPX) != 0 {
				names[i] = t.Field(i).Name

				updateMarketplaceItem := bson.M{
					strings.ToLower(names[i]): pX.Interface(),
				}
				result, err := models.UpdateMarketplaceItemWithID(objItemId, updateMarketplaceItem, c)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in updating marketplace item with ID."})
					return
				}
				fmt.Printf("Update Result of Marketpace Item => %v", result)
			}
		}
		updatedMarketplaceItem, err := models.GetMarketplaceItemOne(objItemId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred getting updated marketplace item."})
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedMarketplaceItem})
	}
}

func DeleteMarketplaceWithId() gin.HandlerFunc {
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
		marketplaceId := ctx.Param("marketplaceId")

		defer cancel()

		objMarketplaceId, err := strconv.Atoi(marketplaceId)

		result, err := models.DeleteMarketplaceOne(objMarketplaceId, c)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}
		fmt.Printf("Result of delete marketplace type => %v", result)
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": "Marketplace Type Successfully Deleted."})
	}
}

func DeleteMarketplaceItem() gin.HandlerFunc {
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
		marketplaceItemId := ctx.Param("itemId")
		defer cancel()

		objMarketplaceItemId, _ := primitive.ObjectIDFromHex(marketplaceItemId)

		result, err := models.DeleteMarketplaceItemOne(objMarketplaceItemId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}
		fmt.Printf("Marketplace Item Delete Result => %v", result)
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": "Marketplace Item successfully deleted."})
	}
}
