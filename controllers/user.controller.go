package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"serendipity_backend/SerendipityRequest"
	"serendipity_backend/SerendipityResponse"

	"serendipity_backend/configs"
	"serendipity_backend/models"
	"serendipity_backend/utilities"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func unique(s []primitive.ObjectID) []primitive.ObjectID {
	inResult := make(map[primitive.ObjectID]bool)
	var result []primitive.ObjectID
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func SignInUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input SerendipityRequest.LoginEmailRequest
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "JSON Binding Error."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Eccor Occurred in validating request body."})
			return
		}

		user, err := models.LoginCheck(input.Email, input.Password, c)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Checking Login failed"})
			return
		}

		config, _ := configs.LoadConfig(".")
		access_token, err := utilities.CreateToken(config.AccessTokenExpiresIn, user.Email, user.Role, config.AccessTokenPrivateKey)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured for getting user token.", "status": "failed"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": gin.H{"token": access_token, "data": user}})
	}
}

func CreateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input SerendipityRequest.SignUpEmailRequest
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Eccor Occured in binding json."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Eccor Occurred in validating request body."})
			return
		}

		if _, err := models.ValidateEmail(input.Email); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error Occurred in validating email."})
			return
		}

		hashedPassword, _ := utilities.HashPassword(input.Password)
		newUser := models.User{
			ID:          primitive.NewObjectID(),
			FirstName:   input.FirstName,
			LastName:    input.LastName,
			Email:       input.Email,
			Password:    hashedPassword,
			PhoneNumber: input.PhoneNumber,
			Activate:    true,
			Role:        3,
			LoginStatus: true,
		}

		curUser, err := newUser.SaveUser(c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error Occurred for add a new user."})
			return
		}

		config, _ := configs.LoadConfig(".")
		access_token, err := utilities.CreateToken(config.AccessTokenExpiresIn, input.Email, 3, config.AccessTokenPrivateKey)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error Occurred in getting user token."})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": gin.H{"data": curUser, "token": access_token}})
	}
}

func GoogleAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input SerendipityRequest.GoogleAuthRequest
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "JSON Binding Error."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Eccor Occurred in validating request body."})
			return
		}

		user, err := models.GoogleAuthCheck(input.Email, c)
		if err != nil {
			newUser := models.User{
				ID:         primitive.NewObjectID(),
				FirstName:  input.FirstName,
				LastName:   input.LastName,
				Email:      input.Email,
				Avatar:     input.Avatar,
				SocialType: input.SocialType,
				SocialId:   input.SocialId,
				Role:       3,
			}
			nUser, err := newUser.SaveUser(c)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error Occurred for add a new user."})
				return
			}

			config, _ := configs.LoadConfig(".")
			access_token, err := utilities.CreateToken(config.AccessTokenExpiresIn, input.Email, 3, config.AccessTokenPrivateKey)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error Occurred in getting user token."})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": gin.H{"data": nUser, "token": access_token}})
		} else {
			config, _ := configs.LoadConfig(".")
			access_token, err := utilities.CreateToken(config.AccessTokenExpiresIn, user.Email, user.Role, config.AccessTokenPrivateKey)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error Occured for getting user token.", "status": "failed"})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": gin.H{"token": access_token, "data": user}})
		}
	}
}

func UpdateUserInfo() gin.HandlerFunc {
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
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := ctx.Param("userId")
		var user SerendipityRequest.UpdateProfileRequest
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		// validate the request body
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error Occurred in Binding JSON."})
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error Occurred in validating request body."})
			return
		}

		selectedUser, err := models.FindUserWithID(objId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}
		fmt.Printf("Selected User With ID => %v", selectedUser)

		t := reflect.TypeOf(user)
		names := make([]string, t.NumField())
		for i := range names {
			pX := getAttr(&user, t.Field(i).Name)
			fmtedPX := fmt.Sprintf("%v", pX)
			if (len(fmtedPX) != 0) && (len(fmtedPX) != 2) {
				names[i] = t.Field(i).Name

				var updateUserInfo bson.M

				if strings.ToLower(names[i]) == "password" {
					hashedPassword, _ := utilities.HashPassword(user.Password)
					updateUserInfo = bson.M{
						"password": hashedPassword,
					}
				} else {
					updateUserInfo = bson.M{
						strings.ToLower(names[i]): pX.Interface(),
					}
				}
				result, err := models.UpdateUser(c, updateUserInfo, objId)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": SerendipityResponse.UPDATE_USER_ERROR})
					return
				}
				fmt.Printf("Update Result of user info => %v", result)
			}
		}
		updatedUserInfo, err := models.FindUserWithID(objId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred getting updated user."})
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedUserInfo})
	}
}

func UpdateProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "User is not lgged in."})
			return
		}
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := ctx.Param("userId")
		var user SerendipityRequest.UpdateProfileRequest
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		// validate the request body
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error Occurred in Binding JSON."})
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error Occurred in validating request body."})
			return
		}

		selectedUser, err := models.FindUpdateUserWithIDAndEmail(objId, user_email, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}
		fmt.Printf("Selected User With ID => %v", selectedUser)

		t := reflect.TypeOf(user)
		names := make([]string, t.NumField())
		for i := range names {
			pX := getAttr(&user, t.Field(i).Name)
			fmtedPX := fmt.Sprintf("%v", pX)
			if (len(fmtedPX) != 0) && (len(fmtedPX) != 2) {
				names[i] = t.Field(i).Name

				var updateUserInfo bson.M

				if strings.ToLower(names[i]) == "password" {
					hashedPassword, _ := utilities.HashPassword(user.Password)
					updateUserInfo = bson.M{
						"password": hashedPassword,
					}
				} else {
					updateUserInfo = bson.M{
						strings.ToLower(names[i]): pX.Interface(),
					}
				}
				result, err := models.UpdateUserProfile(c, updateUserInfo, objId, user_email)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": SerendipityResponse.UPDATE_USER_ERROR})
					return
				}
				fmt.Printf("Update Result of user info => %v", result)
			}
		}
		updatedUserInfo, err := models.FindUpdateUserWithIDAndEmail(objId, user_email, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred getting updated user."})
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedUserInfo})
	}
}

func UpdateUserRole() gin.HandlerFunc {
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
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := ctx.Param("userId")
		var userRole SerendipityRequest.UpdateUserRole
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		// validate the request body
		if err := ctx.BindJSON(&userRole); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error Occurred in Binding JSON."})
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&userRole); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error Occurred in validating request body."})
			return
		}

		selectedUser, err := models.FindUserWithID(objId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}
		fmt.Printf("Selected User With ID => %v", selectedUser)

		updateUserRole := bson.M{
			"role": userRole.Role,
		}

		result, err := models.UpdateUser(c, updateUserRole, objId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": SerendipityResponse.UPDATE_USER_ERROR})
			return
		}
		fmt.Printf("Update Result of user info => %v", result)

		updatedUserInfo, err := models.FindUserWithID(objId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred getting updated user."})
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedUserInfo})
	}
}

func DeleteUser() gin.HandlerFunc {
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
		userId := ctx.Param("userId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := models.DeleteUserOne(objId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error Occurred in deleting User."})
			return
		}
		fmt.Printf("User Delete Result => %v", result)
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": "User Successfully deleted."})
	}
}

func HandleFollow() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := ctx.Param("userId")
		var user SerendipityRequest.FollowForumPostwithUserId
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)
		defer cancel()

		objUserId, _ := primitive.ObjectIDFromHex(userId)

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validating of request."})
			return
		}

		result, err := models.FindUserWithID(objUserId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		result.Follows = append(result.Follows, user.UserID)

		result.Follows = unique(result.Follows)

		updateUserFollows := bson.M{
			"follows": result.Follows,
		}

		updateResult, err := models.UpdateUser(c, updateUserFollows, objUserId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		fmt.Printf("Updated User Result => %v", updateResult)

		updatedUser, err := models.FindUserWithID(objUserId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedUser})
	}
}

func SetUserFoundation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "failed"})
			return
		}
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for Update is %v", user_email)

		userId := ctx.Param("userId")
		var foundation SerendipityRequest.SetUserFoundation
		defer cancel()

		objUserId, _ := primitive.ObjectIDFromHex(userId)

		if er := ctx.ShouldBindJSON(&foundation); er != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": er.Error(), "status": "Error occurred in JSON Binding."})
			return
		}

		if validationErr := validate.Struct(&foundation); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validating of request."})
			return
		}

		selectedUser, err := models.FindUserWithID(objUserId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		fmt.Printf("Selected User With ID => %v", selectedUser)

		updateUserFoundation := bson.M{
			"foundation": foundation.Foundation,
		}

		result, errr := models.UpdateUser(c, updateUserFoundation, objUserId)
		if errr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}
		fmt.Printf("Update Result of User Foundation => %v", result)

		selectedFoundation, errrr := models.GetFoundationWithID(foundation.Foundation, c)
		if errrr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "userFoundation": selectedFoundation})
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")
		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "User is not lgged in."})
			return
		}
		role, exists := ctx.Get("role")
		user_role, _ := strconv.Atoi(fmt.Sprintf("%v", role))
		fmt.Printf("User Role for Update => %v", user_role)

		// if user_role == 3 {
		// 	ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_PERMISSION_ALLOWED, "status": "Permission is not arranged."})
		// 	return
		// }
		user_email := fmt.Sprintf("User Email %v", email)
		fmt.Printf("User Email => %v", user_email)
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		limit := ctx.Request.URL.Query().Get("results")
		page := ctx.Request.URL.Query().Get("page")
		sortField := ctx.Request.URL.Query().Get("sortField")
		sortOrder := ctx.Request.URL.Query().Get("sortOrder")

		convertedLimit, er := strconv.Atoi(limit)
		if er != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": er.Error(), "status": "Limit Atoi failed"})
			return
		}

		convertedPage, err := strconv.Atoi(page)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Page Atoi failed"})
			return
		}

		if user_role == 1 {
			AdminUserResults, err := models.FindAllAdminsUsers(convertedLimit, convertedPage, sortField, sortOrder, 1, c)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": AdminUserResults})
		} else if (user_role == 2) || (user_role == 3) {
			UserResults, err := models.FindAllUsers(convertedLimit, convertedPage, sortField, sortOrder, 2, c)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": UserResults})
		}
	}
}

func UpdateUserActivation() gin.HandlerFunc {
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
		var input SerendipityRequest.UpdateUserActivation
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "JSON Binding Error."})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error Occurred in validating request body."})
			return
		}

		userId := ctx.Param("userId")
		objUserId, _ := primitive.ObjectIDFromHex(userId)

		selectedUser, err := models.FindUserWithID(objUserId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "failed"})
			return
		}
		fmt.Printf("Selected User With ID => %v", selectedUser)

		updateUserInfo := bson.M{
			"activate": input.Activate,
		}

		result, err := models.UpdateUser(c, updateUserInfo, objUserId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": SerendipityResponse.UPDATE_USER_ERROR})
			return
		}

		fmt.Printf("Update Result of User Info => %v", result)

		updatedUserInfo, err := models.FindUserWithID(objUserId, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred getting updated user."})
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedUserInfo})
	}
}

func LogOut() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email, exists := ctx.Get("email")

		if !exists {
			ctx.JSON(http.StatusBadGateway, gin.H{"error": SerendipityResponse.NOT_LOGGED_IN, "status": "User is not logged in."})
			return
		}
		user_email := fmt.Sprint(email)
		fmt.Printf("User Email for log out is %v", user_email)
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		result, err := models.SetLogOut(c, user_email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": SerendipityResponse.ERROR_UPDATE_USER_LOGINSTATUS})
			return
		}
		fmt.Printf("Update Result of User Info => %v", result)
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": "User has been logged out successfully!"})
	}
}

func SetLogoutWithID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input SerendipityRequest.SetLogoutWithID
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "JSON Binding Error Occurred"})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error Occurred in validating request body."})
			return
		}

		selectedUser, err := models.FindUserWithID(input.UserID, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Failed"})
			return
		}
		fmt.Printf("Selected User With ID => %v", selectedUser)

		updateUserInfo := bson.M{
			"loginstatus": false,
		}

		result, err := models.UpdateUser(c, updateUserInfo, input.UserID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": SerendipityResponse.UPDATE_USER_ERROR})
			return
		}
		fmt.Printf("Update Result of User Info => %v", result)

		updatedUserInfo, err := models.FindUserWithID(input.UserID, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error Occurred Getting Updated User."})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedUserInfo})
	}
}

func CheckUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input SerendipityRequest.CheckUserRequest
		defer cancel()

		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "JSON Binding "})
			return
		}

		if validationErr := validate.Struct(&input); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error Occurred in validating request body."})
			return
		}

		user, err := models.GetUserWithEmail(c, input.Email)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Checking User failed"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": user})
	}
}

func SendMail() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		generated_random_code := models.EncodeToString(6)
		get_current_timestamp := time.Now().Unix()

		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user SerendipityRequest.PasswordRestRequest
		defer cancel()

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error Occurred in Binding JSON."})
			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error Occurred in validating structure."})
			return
		}

		var updateUserInfo bson.M

		updateUserInfo = bson.M{
			"verificationcode": generated_random_code,
			"expiration":       get_current_timestamp,
		}

		result, err := models.UpdateUserWithEmail(c, updateUserInfo, user.Email)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": SerendipityResponse.UPDATE_USER_ERROR})
			return
		}

		fmt.Printf("Updated Result of User Info for password reset => %v", result)

		updatedUserInfo, err := models.FindUserWithEmail(user.Email, c)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": "Error occurred in getting user with email."})
			return
		}

		fmt.Printf("Generated Random Code => %v", generated_random_code)
		fmt.Printf("Current TimeStamp => %v", get_current_timestamp)

		news := []models.Article{
			{
				Image:   "https://ourserendipity.s3.amazonaws.com/logo.png",
				Content: generated_random_code,
			},
		}
		cwd, _ := os.Getwd()
		// var errr error
		// Get newsletter html template in string format
		htmlContent, err := ParseTemplate(filepath.Join(cwd, "templates", "./newsletter.html"), news)
		if err != nil {
			log.Fatalln(err)
		}

		//Get newsletter text template in string format
		textContent, err := ParseTemplate(filepath.Join(cwd, "templates", "./newsletter.tmpl"), news)
		if err != nil {
			log.Fatalln(err)
		}

		mailRequest := NewMailRequest(
			"no-reply@golang.coach",
			"Serendipity Forgot Password",
			htmlContent,
			textContent,
			[]string{user.Email},
		)

		// send mail
		ok, err := mailRequest.SendMail()
		if !ok {
			log.Fatal(err)
		}

		fmt.Println("Mail has been sent.")

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": updatedUserInfo})
	}
}

func ResetCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user SerendipityRequest.PasswordRestRequest
		defer cancel()

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in binding JSON."})
			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validating structure."})
			return
		}

		var updateUserInfo bson.M

		updateUserInfo = bson.M{
			"verificationcode": "",
			"expiration":       0,
		}

		result, err := models.UpdateUserWithEmail(c, updateUserInfo, user.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": SerendipityResponse.UPDATE_USER_ERROR})
			return
		}
		fmt.Printf("Updated result for reset of verification code => %v", result)

		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": "Successfully reseted for this user's vericication code and expiration !"})
	}
}

func ResetPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user SerendipityRequest.ResetWithNewPasswordRequest
		defer cancel()

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "Error occurred in binding json"})
			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error(), "status": "Error occurred in validation JSON structure."})
			return
		}

		var updateUserPassword bson.M
		hashedPassword, _ := utilities.HashPassword(user.NewPassword)
		updateUserPassword = bson.M{
			"password": hashedPassword,
		}

		result, err := models.UpdateUserWithEmail(c, updateUserPassword, user.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "status": SerendipityResponse.UPDATE_USER_ERROR})
			return
		}
		fmt.Printf("Updated Result for password reset => %v", result)

		ctx.JSON(http.StatusOK, gin.H{"status": true, "payload": "Password has been updated successfully !"})
	}
}
