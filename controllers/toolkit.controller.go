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

var toolkitTypeCollection *mongo.Collection = configs.GetCollection(configs.DB, "toolkit")
var toolkitPostCollection *mongo.Collection = configs.GetCollection(configs.DB, "toolkitPost")

func getAttr(obj interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		panic("not struct")
	}
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		panic("not found:" + fieldName)
	}
	return curField
}

func AddNewToolkit() gin.HandlerFunc {
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
		var input SerendipityRequest.AddToolkitRequest
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error Occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error Occurred in validation of request."})
			return
		}

		newToolkitType := models.ToolkitType{
			ID:               primitive.NewObjectID(),
			Title:            input.Title,
			CoverLetterImage: input.CoverLetterImage,
			SortType:         input.SortType,
			Type:             input.Type,
		}

		curToolkit, err := newToolkitType.SaveToolkit(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in creating a new toolkit."})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": curToolkit})
	}
}

func AddNewToolkitPost() gin.HandlerFunc {
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
		currentTime := time.Now()
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input SerendipityRequest.AddToolkitPostRequest
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validation of request."})
			return
		}

		newToolkitPost := models.Post{
			ID:               primitive.NewObjectID(),
			ToolkitType:      input.ToolkitType,
			Title:            input.Title,
			Description:      input.Description,
			CoverLetterImage: input.CoverLetterImage,
			CookingPeriod:    input.CookingPeriod,
			Preparation:      input.Preparation,
			Ingredients:      input.Ingredients,
			Instructions:     input.Instructions,
			Medias:           input.Medias,
			SortTypeId:       input.SortTypeId,
			Link:             input.Link,
			PostedAt:         currentTime.Format("2006-01-02 15:04:05"),
		}

		curToolkitPost, err := newToolkitPost.SaveToolkitPost(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in creating a new toolkit post."})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": curToolkitPost})
	}
}

func UpdateToolkit() gin.HandlerFunc {
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
		toolkitId := ctx.Param("toolkitId")
		var toolkitType SerendipityRequest.UpdateToolkitRequest
		defer cancel()

		objToolkitId, _ := primitive.ObjectIDFromHex(toolkitId)

		if err := ctx.ShouldBindJSON(&toolkitType); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&toolkitType); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validation of request."})
			return
		}

		result, err := models.GetToolkitTypeWithID(objToolkitId, c)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Could not get toolkit type with id."})
		}
		if result.ID.Hex() == "000000000000000000000000" {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "The Toolkit is not existed"})
			return
		}

		t := reflect.TypeOf(toolkitType)
		names := make([]string, t.NumField())
		for i := range names {
			pX := getAttr(&toolkitType, t.Field(i).Name)
			fmtedPX := fmt.Sprintf("%v", pX)
			if (len(fmtedPX) != 0) && (len(fmtedPX) != 2) {
				names[i] = t.Field(i).Name
				updateToolkitType := bson.M{
					strings.ToLower(names[i]): pX.Interface(),
				}
				result, err := models.UpdateToolkitType(objToolkitId, updateToolkitType, c)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in updating toolkit type."})
					return
				}
				fmt.Printf("Result of Update Toolkit Type => %v", result)
			}
		}
		updatedToolkitType, errr := models.GetToolkitTypeWithID(objToolkitId, c)
		if errr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Could not get updated toolkit type with id."})
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedToolkitType})
	}
}

func UpdateToolkitPost() gin.HandlerFunc {
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
		postId := ctx.Param("postId")
		var toolkitPost SerendipityRequest.UpdateToolkitPostRequest
		defer cancel()

		objPostId, _ := primitive.ObjectIDFromHex(postId)

		if err := ctx.ShouldBindJSON(&toolkitPost); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&toolkitPost); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validating of request."})
			return
		}

		getToolkitPost, err := models.GetToolkitPostOne(objPostId, c)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Could not get Toolkit Post."})
			return
		}

		fmt.Printf("Get Toolkit Post Result => %v", getToolkitPost)

		t := reflect.TypeOf(toolkitPost)
		names := make([]string, t.NumField())
		for i := range names {
			pX := getAttr(&toolkitPost, t.Field(i).Name)
			fmtedPX := fmt.Sprintf("%v", pX)
			fmt.Printf("Attribute => %v", pX.Interface())
			fmt.Printf("Attribute 2 => %v", len(fmtedPX))
			if pX.Interface() != 0 && len(fmtedPX) != 0 && len(fmtedPX) != 2 {
				names[i] = t.Field(i).Name

				updateToolkitPost := bson.M{
					strings.ToLower(names[i]): pX.Interface(),
				}
				result, err := models.UpdateToolkitPostWithID(objPostId, updateToolkitPost, c)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in updating toolkit post with ID."})
					return
				}
				fmt.Printf("Update Result of Toolkist Post => %v", result)
			}
		}
		updatedToolkitPost, err := models.GetToolkitPostOne(objPostId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred getting updated toolkit Post."})
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedToolkitPost})
	}
}

func SetToolkitWeeklyPost() gin.HandlerFunc {
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
		postId := ctx.Param("postId")
		var weeklyToolkitPost SerendipityRequest.SetWeeklyPost
		defer cancel()

		objPostId, _ := primitive.ObjectIDFromHex(postId)

		if err := ctx.ShouldBindJSON(&weeklyToolkitPost); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&weeklyToolkitPost); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validating of request."})
			return
		}

		getToolkitPost, err := models.GetToolkitPostOne(objPostId, c)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Could not get Toolkit Post."})
			return
		}

		fmt.Printf("Get Toolkit Post Result => %v", getToolkitPost)

		updateWeeklyToolkitPost := bson.M{
			"weeklypost": weeklyToolkitPost.WeeklyPost,
		}

		result, err := models.UpdateToolkitPostWithID(objPostId, updateWeeklyToolkitPost, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in updating toolkit post with ID."})
			return
		}
		fmt.Printf("Update Result of Toolkist Post => %v", result)
		updatedToolkitPost, err := models.GetToolkitPostOne(objPostId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred getting updated toolkit Post."})
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedToolkitPost})
	}
}

func GetAllToolkitTypes() gin.HandlerFunc {
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

		// limit := ctx.Request.URL.Query().Get("results")
		// page := ctx.Request.URL.Query().Get("page")
		// sortField := ctx.Request.URL.Query().Get("sortField")
		// sortOrder := ctx.Request.URL.Query().Get("sortOrder")

		// convertedLimit, er := strconv.Atoi(limit)
		// if er != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": er.Error(), "status": "failed"})
		// 	return
		// }
		// convertedPage, err := strconv.Atoi(page)
		// if err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "failed"})
		// 	return
		// }

		// results, err := models.FindAllToolkitTypes(convertedLimit, convertedPage, sortField, sortOrder, c)
		results, err := models.FindAllToolkitTypes(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": results})
	}
}

func GetAllPostsInToolkit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		toolkitId := ctx.Param("toolkitType")
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		defer cancel()

		objToolkitType, err := strconv.Atoi(toolkitId)
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
		result, errr := models.GetAllToolkitPostsInToolkit(convertedLimit, convertedPage, sortField, sortOrder, objToolkitType, c)

		if errr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errr.Error(), "status": "Error occurred in getting toolkit posts."})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": result})
	}
}

func SetTodayActivities() gin.HandlerFunc {
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
		var todayPost SerendipityRequest.SetTodayActivity
		// var toolkitPosts []models.Post
		defer cancel()

		if err := ctx.ShouldBindJSON(&todayPost); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&todayPost); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validating of request."})
			return
		}

		// results, err := models.GetAllPostsInToolkit(c)
		// if err != nil {
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
		// 	return
		// }

		// for k := 0; k < len(todayPosts.TodayActivities); k++ {
		updatedToolkitPostWithTodayActivity := bson.M{
			"todayactivity": todayPost.ActivityStatus,
		}

		result, err := models.UpdateToolkitPostWithID(todayPost.ToolkitPostID, updatedToolkitPostWithTodayActivity, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}
		fmt.Printf("Updated Result of Toolkist Post => %v", result)
		// }

		// var filteredToolkitPosts []models.Post
		// for i := 0; i < len(results); i++ {
		// 	var count int = 0
		// 	for k := 0; k < len(todayPosts.TodayActivities); k++ {
		// 		if results[i].ID != todayPosts.TodayActivities[k] {
		// 			count++
		// 		}
		// 	}
		// 	if count == len(todayPosts.TodayActivities) {
		// 		filteredToolkitPosts = append(filteredToolkitPosts, results[i])
		// 	}
		// }

		// for i := 0; i < len(filteredToolkitPosts); i++ {
		// 	updatedToolkitPostWithTodayActivity := bson.M{
		// 		"todayactivity": false,
		// 	}
		// 	result, err := models.UpdateToolkitPostWithID(filteredToolkitPosts[i].ID, updatedToolkitPostWithTodayActivity, c)
		// 	if err != nil {
		// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
		// 		return
		// 	}
		// 	fmt.Printf("Filtered Toolkit Posts => %v", result)
		// }

		updateResult, err := models.GetToolkitPostOne(todayPost.ToolkitPostID, c)

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": gin.H{"data": updateResult, "result": "Successfully set up for today's activity."}})
	}
}

func GetToolkitPostsForToday() gin.HandlerFunc {
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

		results, err := models.GetAllPostsInToolkit(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		var todayToolkitPosts []models.Post
		for i := 0; i < len(results); i++ {
			if results[i].TodayActivity == true {
				todayToolkitPosts = append(todayToolkitPosts, results[i])
			}
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": todayToolkitPosts})
	}
}

func DeleteToolkitWithId() gin.HandlerFunc {
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
		toolkitId := ctx.Param("toolkitId")

		defer cancel()

		objToolkitId, err := strconv.Atoi(toolkitId)

		result, err := models.DeleteToolkitOne(objToolkitId, c)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}
		fmt.Printf("Result of delete toolkit type => %v", result)
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": "Toolkit Type Successfully Deleted."})
	}
}

func DeleteToolkitPost() gin.HandlerFunc {
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
		toolkitPostId := ctx.Param("postId")
		defer cancel()

		objPostId, _ := primitive.ObjectIDFromHex(toolkitPostId)

		result, err := models.DeleteToolkitPostOne(objPostId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}
		fmt.Printf("Toolkit Post Delete Result => %v", result)
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": "Toolkit Post successfully deleted."})
	}
}

func GetWeeklyPosts() gin.HandlerFunc {
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

		results, err := models.GetAllWeeklyPostsInToolkit(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": results})
	}
}
