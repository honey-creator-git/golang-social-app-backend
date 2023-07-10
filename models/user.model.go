package models

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"net/mail"
	"serendipity_backend/SerendipityResponse"
	"serendipity_backend/utilities"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Article struct {
	Image   string
	Content string
}

type User struct {
	ID               primitive.ObjectID   `json:"id,omitempty"`
	Email            string               `json:"email,omitempty"`
	FirstName        string               `json:"firstName,omitempty"`
	LastName         string               `json:"lastName,omitempty"`
	Password         string               `json:"password,omitempty"`
	PhoneNumber      string               `json:"phoneNumber,omitempty"`
	SocialType       string               `json:"socialType"`
	SocialId         string               `json:"socialId"`
	PushNotification bool                 `json:"pushNotification"`
	Avatar           string               `json:"avatar"`
	Follows          []primitive.ObjectID `json:"follows"`
	Foundation       string               `json:"foundation"`
	Role             int                  `json:"role,omitempty"`     // 1 - Super Admin / 2 - Admin / 3 - User
	Activate         bool                 `json:"activate,omitempty"` // true, false
	LoginStatus      bool                 `json:"loginStatus,omitempty"`
	VerificationCode string               `json:"verificationCode"`
	Expiration       int                  `json:"expiration"`
}

func validMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}

func ValidateEmail(email string) (bool, error) {
	add, ok := validMailAddress(email)
	if !ok {
		return false, errors.New(SerendipityResponse.EMAIL_VERIFY_ERROR)
	} else {
		fmt.Printf("Email is %v", add)
		return true, nil
	}
}

func LoginCheck(email string, password string, c context.Context) (User, error) {
	var err error

	user := User{}

	updateUserInfo := bson.M{
		"loginstatus": true,
	}

	err = userCollection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return User{}, errors.New(SerendipityResponse.EMAIL_NOT_FOUND)
	}

	err = utilities.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return user, errors.New(SerendipityResponse.WRONG_PASSWORD)
	}

	result, err := userCollection.UpdateOne(c, bson.M{"email": email}, bson.M{"$set": updateUserInfo})
	fmt.Printf("Result of update of user's login status => %v", result)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return User{}, errors.New("Email Already Exists")
		}
		return User{}, errors.New(SerendipityResponse.UPDATE_USER_ERROR)
	}

	return user, nil
}

func (newUser *User) SaveUser(c context.Context) (User, error) {

	mUser, err := userCollection.InsertOne(c, newUser)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return User{}, errors.New("Email Already Exists")
		}
	}

	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := userCollection.Indexes().CreateOne(c, index); err != nil {
		return User{}, errors.New("could not create index of email")
	}

	var curUser User
	userCollection.FindOne(c, bson.M{"_id": mUser.InsertedID}).Decode(&curUser)

	return curUser, nil
}

func GoogleAuthCheck(email string, c context.Context) (User, error) {
	var err error

	user := User{}
	err = userCollection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return User{}, errors.New(SerendipityResponse.EMAIL_NOT_FOUND)
	}

	return user, nil
}

func UpdateUser(c context.Context, payload map[string]interface{}, id primitive.ObjectID) (User, error) {

	result, err := userCollection.UpdateOne(c, bson.M{"id": id}, bson.M{"$set": payload})
	fmt.Printf("Updated User Result Count => %v", result)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return User{}, errors.New("Email Already Exists")
		}
		return User{}, errors.New(SerendipityResponse.UPDATE_USER_ERROR)
	}
	if result.MatchedCount < 1 {
		return User{}, errors.New(SerendipityResponse.ERROR_UPDATE_USER)
	}
	return User{}, nil
}

func UpdateUserProfile(c context.Context, payload map[string]interface{}, id primitive.ObjectID, email string) (User, error) {

	result, err := userCollection.UpdateOne(c, bson.M{"id": id, "email": email}, bson.M{"$set": payload})
	fmt.Printf("Updated User Profile => %v", result)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return User{}, errors.New("Email Already Exists")
		}
		return User{}, errors.New(SerendipityResponse.UPDATE_USER_ERROR)
	}
	if result.MatchedCount < 1 {
		return User{}, errors.New(SerendipityResponse.ERROR_UPDATE_USER)
	}
	return User{}, nil
}

func DeleteUserOne(objId primitive.ObjectID, c context.Context) (User, error) {
	result, err := userCollection.DeleteOne(c, bson.M{"id": objId})
	fmt.Printf("Deleted User Result Count => %v", result.DeletedCount)
	if err != nil {
		return User{}, errors.New(SerendipityResponse.DELETE_USER_ERROR)
	}
	if result.DeletedCount < 1 {
		return User{}, errors.New(SerendipityResponse.NOT_FOUND_USER_DELETE)
	}
	return User{}, nil
}

func FindUserWithID(objId primitive.ObjectID, c context.Context) (User, error) {
	var user User
	userCollection.FindOne(c, bson.M{"id": objId}).Decode(&user)

	if user.ID.Hex() == "000000000000000000000000" {
		return User{}, errors.New(SerendipityResponse.NOT_EXIST_USER)
	}

	return user, nil
}

func FindUserWithEmail(email string, c context.Context) (User, error) {
	var user User
	userCollection.FindOne(c, bson.M{"email": email}).Decode(&user)

	if user.ID.Hex() == "000000000000000000000000" {
		return User{}, errors.New(SerendipityResponse.NOT_EXIST_USER)
	}

	return user, nil
}

func FindUpdateUserWithIDAndEmail(objId primitive.ObjectID, email string, c context.Context) (User, error) {
	var user User
	userCollection.FindOne(c, bson.M{"id": objId, "email": email}).Decode(&user)

	if user.ID.Hex() == "000000000000000000000000" {
		return User{}, errors.New(SerendipityResponse.NOT_EXIST_USER_UPDATE_PROFILE)
	}

	return user, nil
}

func FindAllUsers(limit int, page int, sortField string, sortOrder string, role int, c context.Context) ([]User, error) {
	skipNumber := (page - 1) * limit
	pageOptions := options.Find()
	if sortOrder == "ascend" {
		pageOptions.SetSort(bson.D{primitive.E{Key: sortField, Value: 1}})
	} else if sortOrder == "descend" {
		pageOptions.SetSort(bson.D{primitive.E{Key: sortField, Value: -1}})
	}
	pageOptions.SetSkip(int64(skipNumber))
	pageOptions.SetLimit(int64(limit))

	var users []User
	results, err := userCollection.Find(c, bson.M{"role": bson.M{"$gt": role}}, pageOptions)
	if err != nil {
		return []User{}, errors.New(SerendipityResponse.ERROR_GET_USERS)
	}

	defer results.Close(c)

	for results.Next(c) {
		var singleUser User
		if err := results.Decode(&singleUser); err != nil {
			return []User{}, errors.New(SerendipityResponse.ERROR_DECODE_USER)
		}
		users = append(users, singleUser)
	}

	return users, nil
}

func FindAllAdminsUsers(limit int, page int, sortField string, sortOrder string, role int, c context.Context) ([]User, error) {
	skipNumber := (page - 1) * limit
	pageOptions := options.Find()
	if sortOrder == "ascend" {
		pageOptions.SetSort(bson.D{primitive.E{Key: sortField, Value: 1}})
	} else if sortOrder == "descend" {
		pageOptions.SetSort(bson.D{primitive.E{Key: sortField, Value: -1}})
	}
	pageOptions.SetSkip(int64(skipNumber))
	pageOptions.SetLimit(int64(limit))

	var admins []User
	results, err := userCollection.Find(c, bson.M{"role": bson.M{"$gt": role}}, pageOptions)
	if err != nil {
		return []User{}, errors.New(SerendipityResponse.ERROR_GET_ADMINS)
	}

	defer results.Close(c)

	for results.Next(c) {
		var singleAdmin User
		if err := results.Decode(&singleAdmin); err != nil {
			return []User{}, errors.New(SerendipityResponse.ERROR_DECODE_USER)
		}
		admins = append(admins, singleAdmin)
	}

	return admins, nil
}

func SetLogOut(c context.Context, email string) (User, error) {
	updateUserInfo := bson.M{
		"loginstatus": false,
	}
	result, err := userCollection.UpdateOne(c, bson.M{"email": email}, bson.M{"$set": updateUserInfo})
	fmt.Printf("Result of Log Out for User => %v", result)

	if err != nil {
		return User{}, errors.New(SerendipityResponse.ERROR_UPDATE_USER_LOGINSTATUS)
	}
	if result.MatchedCount < 1 {
		return User{}, errors.New(SerendipityResponse.ERROR_UPDATE_USER)
	}
	return User{}, nil
}

func GetUserWithEmail(c context.Context, email string) (User, error) {
	var err error

	user := User{}

	err = userCollection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return User{}, errors.New(SerendipityResponse.EMAIL_NOT_FOUND)
	}

	return user, nil
}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func UpdateUserWithEmail(c context.Context, payload map[string]interface{}, email string) (User, error) {

	result, err := userCollection.UpdateOne(c, bson.M{"email": email}, bson.M{"$set": payload})
	fmt.Printf("Updated for password reset => %v", result)
	if err != nil {
		return User{}, errors.New(SerendipityResponse.UPDATE_USER_ERROR)
	}
	if result.MatchedCount < 1 {
		return User{}, errors.New(SerendipityResponse.ERROR_UPDATE_USER)
	}
	return User{}, nil
}
