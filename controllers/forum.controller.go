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

var forumPostCollection *mongo.Collection = configs.GetCollection(configs.DB, "forumPost")
var forumTypeCollection *mongo.Collection = configs.GetCollection(configs.DB, "forumType")

func AddForumType() gin.HandlerFunc {
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
		var input SerendipityRequest.AddForumTypeRequest
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validation of request."})
			return
		}

		newForumType := models.ForumType{
			ID:               primitive.NewObjectID(),
			Title:            input.Title,
			CoverLetterImage: input.CoverLetterImage,
			ForumType:        input.ForumType,
		}

		curForumType, err := newForumType.SaveForumType(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in creating a new Forum Type."})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": curForumType})
	}
}

func AddForumPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		currentTime := time.Now()
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		var input SerendipityRequest.AddForumPostRequest
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validation of request."})
			return
		}

		newForumPost := models.ForumPost{
			ID:               primitive.NewObjectID(),
			ForumType:        input.ForumType,
			Title:            input.Title,
			CoverLetterImage: input.CoverLetterImage,
			Description:      input.Description,
			CreatedAt:        currentTime.Format("2006-01-02 15:04:05"),
			CreatedBy:        input.CreatedBy,
			Comments:         input.Comments,
			VisitCount:       input.VisitCount,
			Emotions:         input.Emotions,
		}

		curForumPost, err := newForumPost.SaveForumPost(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in creating a new forum post."})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": curForumPost})
	}
}

func AddCommentForForumPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		currentTime := time.Now()
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		postId := ctx.Param("postId")
		var comment SerendipityRequest.NewCommentForForumPost
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		defer cancel()

		objPostId, _ := primitive.ObjectIDFromHex(postId)

		if err := ctx.ShouldBindJSON(&comment); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&comment); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validation of request."})
			return
		}

		var getForumPost models.ForumPost
		forumPostCollection.FindOne(c, bson.M{"id": objPostId}).Decode(&getForumPost)
		if getForumPost.ID.Hex() == "000000000000000000000000" {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "Could not find the Forum Post"})
			return
		}

		newComment := models.Comment{
			ID:          primitive.NewObjectID(),
			Description: comment.Description,
			PostedAt:    currentTime.Format("2006-01-02 15:04:05"),
			PostedBy:    comment.PostedBy,
			PostId:      comment.PostId,
			Emotions:    comment.Emotions,
		}

		getForumPost.Comments = append(getForumPost.Comments, newComment)

		fmt.Printf("Added Comments => %v", getForumPost.Comments)

		addComment := bson.M{
			"comments": getForumPost.Comments,
		}

		result, err := models.AddCommentForForumPost(objPostId, addComment, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Could not add comment for forum post."})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": gin.H{"updated_comments": result, "result": "successfully added comment to Forum Post."}})
	}
}

func AddForumPostCommentEmotion() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input SerendipityRequest.AddEmotionForForumPostComment
		postId := ctx.Param("postId")
		commentId := ctx.Param("commentId")
		defer cancel()

		objPostId, _ := primitive.ObjectIDFromHex(postId)
		objCommentId, _ := primitive.ObjectIDFromHex(commentId)

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in Binding JSON."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validating of request."})
			return
		}

		forumPost, err := models.GetForumPostWithID(objPostId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in getting Forum Post with ID"})
			return
		}

		var updatedForumPostComments []models.Comment
		var forumPostComment models.Comment
		for i := 0; i < len(forumPost.Comments); i++ {
			forumPostComment = forumPost.Comments[i]
			if forumPostComment.ID == objCommentId {
				forumPostComment.Emotions = input.Emotions
			}
			updatedForumPostComments = append(updatedForumPostComments, forumPostComment)
		}

		update := bson.M{
			"comments": updatedForumPostComments,
		}

		result, err := models.UpdateForumPostComments(objPostId, update, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in updating Forum Post Comments."})
			return
		}
		fmt.Printf("Forum Post Update with Comments => %v", result)

		updatedResult, errr := models.GetForumPostWithID(objPostId, c)
		if errr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in getting updated Forum Post."})
			return
		}

		for i := 0; i < len(updatedResult.Comments); i++ {
			if updatedResult.Comments[i].ID == objCommentId {
				ctx.JSON(http.StatusOK, gin.H{"status": true, "data": gin.H{"payload": updatedResult.Comments[i], "result": "successfully updated forum post comment with Emotion."}})
				return
			}
		}
	}
}

func UpdateForumPostWithEmotions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input SerendipityRequest.UpdateForumPostWithEmotions
		postId := ctx.Param("postId")
		defer cancel()

		objPostId, _ := primitive.ObjectIDFromHex(postId)

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurrred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validating of request."})
			return
		}

		update := bson.M{
			"emotions": input.Emotions,
		}

		result, err := models.UpdateForumPostWithPostID(objPostId, update, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		fmt.Printf("Updated Result of Forum Post => %v", result)

		updatedForumPost, errr := models.GetForumPostWithID(objPostId, c)
		if errr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errr.Error(), "status": "Could not get Updated Forum Post with Emotion."})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedForumPost})
	}
}

func UpdateForumPostWithVisitCount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input SerendipityRequest.UpdateForumPostVisitCount
		postId := ctx.Param("postId")
		defer cancel()

		objPostId, _ := primitive.ObjectIDFromHex(postId)

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurrred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validating of request."})
			return
		}

		update := bson.M{
			"visitcount": input.VisitCount,
		}

		result, err := models.UpdateForumPostWithPostID(objPostId, update, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		fmt.Printf("Updated Result of Forum Post => %v", result)

		updatedForumPost, errr := models.GetForumPostWithID(objPostId, c)
		if errr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errr.Error(), "status": "Could not get Updated Forum Post with Emotion."})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedForumPost})
	}
}

func AddForumPostComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update ddd is %v", user_email)
		currentTime := time.Now()
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input SerendipityRequest.AddNewForumPostCommentRequest
		postId := ctx.Param("postId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(postId)

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "JSON Binding InValide"})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Validation Issue"})
			return
		}

		var forumPost models.ForumPost
		err := forumPostCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&forumPost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		newForumPostComment := models.Comment{
			ID:          primitive.NewObjectID(),
			Description: input.Description,
			PostedAt:    currentTime.Format("2006-01-02 15:04:05"),
			PostedBy:    input.PostedBy,
			PostId:      input.PostId,
			Emotions:    input.Emotions,
		}

		fmt.Printf("NewForumPostComment is %v", newForumPostComment)

		forumPost.Comments = append(forumPost.Comments, newForumPostComment)

		update := bson.M{
			"comments": forumPost.Comments,
		}

		result, err := forumPostCollection.UpdateOne(c, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		var updatedForumPost models.ForumPost
		if result.MatchedCount == 1 {
			err := forumPostCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedForumPost)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
				return
			}
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedForumPost})
	}
}

func GetForumPostsWithId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		forumId := ctx.Param("forumId")
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		defer cancel()

		objForumTypeId, err := strconv.Atoi(forumId)

		results, err := models.GetAllForumPostsWithForumID(c)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in getting forum posts."})
			return
		}

		var returnForumPosts []models.ForumPost

		for i := 0; i < len(results); i++ {
			if objForumTypeId == results[i].ForumType {
				returnForumPosts = append(returnForumPosts, results[i])
			}
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": returnForumPosts})
	}
}

func GetForumPostWithPostId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		postId := ctx.Param("postId")
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		defer cancel()

		objForumPostId, _ := primitive.ObjectIDFromHex(postId)

		result, err := models.GetForumPostWithPostID(objForumPostId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": result})
	}
}

func GetAllForums() gin.HandlerFunc {
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

		results, err := models.GetForumTypes(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": results})
	}
}

func DeleteForumPostWithPostId() gin.HandlerFunc {
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

		defer cancel()

		objPostId, _ := primitive.ObjectIDFromHex(postId)

		err := models.DeleteForumPostOne(objPostId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": "Forum Post is deleted successfully."})
	}
}

func UpdateForumPostWithPostIDUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input SerendipityRequest.UpdateForumPostWithPostIDUserID

		postId := ctx.Param("postId")
		userId := ctx.Param("userId")

		defer cancel()

		objPostId, _ := primitive.ObjectIDFromHex(postId)
		objUserId, _ := primitive.ObjectIDFromHex(userId)

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurrred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validating of request."})
			return
		}

		t := reflect.TypeOf(input)
		names := make([]string, t.NumField())
		for i := range names {
			pX := getAttr(&input, t.Field(i).Name)
			fmtedPX := fmt.Sprintf("%v", pX)
			if len(fmtedPX) != 0 {
				names[i] = t.Field(i).Name
				update := bson.M{
					strings.ToLower(names[i]): pX.Interface(),
				}
				result, err := models.UpdateForumPostWithPostIDUserID(objPostId, objUserId, update, c)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
					return
				}
				fmt.Printf("Result of Update Forum Post => %v", result)
			}
		}

		updatedForumPost, errr := models.GetForumPostWithID(objPostId, c)
		if errr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errr.Error(), "status": "Could not get Updated Forum Post."})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedForumPost})
	}
}

func DeleteForumWithId() gin.HandlerFunc {
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
		forumId := ctx.Param("forumId")

		defer cancel()

		objForumId, err := strconv.Atoi(forumId)

		result, err := models.DeleteForumOne(objForumId, c)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		fmt.Printf("Result of delete forum type => %v", result)
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": "Forum Type Successfully Deleted."})
	}
}

func UpdateForumTypeWithID() gin.HandlerFunc {
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
		forumId := ctx.Param("forumId")
		var forumTypeUpdate SerendipityRequest.UpdateForumType
		defer cancel()

		objForumId, _ := primitive.ObjectIDFromHex(forumId)

		if err := ctx.ShouldBindJSON(&forumTypeUpdate); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&forumTypeUpdate); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validating of request body."})
			return
		}

		result, err := models.FindForumTypeWithID(objForumId, c)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Could not get Forum Type with that id."})
			return
		}

		if result.ID.Hex() == "000000000000000000000000" {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "The Forum is not existed."})
			return
		}

		t := reflect.TypeOf(forumTypeUpdate)
		names := make([]string, t.NumField())
		for i := range names {
			pX := getAttr(&forumTypeUpdate, t.Field(i).Name)
			fmtedPX := fmt.Sprintf("%v", pX)
			if len(fmtedPX) != 0 {
				names[i] = t.Field(i).Name
				updateForumType := bson.M{
					strings.ToLower(names[i]): pX.Interface(),
				}
				result, err := models.UpdateForumType(objForumId, updateForumType, c)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in updateing forum type with id."})
					return
				}
				fmt.Printf("Result of Update Forum Type => %v", result)
			}
		}

		updateForumType, err := models.FindForumTypeWithID(objForumId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Could not get updated forum type with id."})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updateForumType})
	}
}
