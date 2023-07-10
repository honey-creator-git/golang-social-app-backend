package models

import (
	"context"
	"errors"
	"serendipity_backend/SerendipityResponse"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Media struct {
	Url       string `json:"url,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
	Period    int    `json:"period,omitempty"`
}

type Post struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	ToolkitType      int                `json:"toolkitType,omitempty"`
	Title            string             `json:"title,omitempty"`
	Description      string             `json:"description"`
	CookingPeriod    int                `json:"cookingPeriod"`
	Preparation      int                `json:"preparation"`
	Ingredients      []string           `json:"ingredients"`
	Instructions     []string           `json:"instructions"`
	CoverLetterImage string             `json:"coverLetterImage"`
	Medias           []Media            `json:"medias"`
	PostedAt         string             `json:"postedAt,omitempty"`
	SortTypeId       string             `json:"sortTypeId"`
	TodayActivity    bool               `json:"todayActivity"`
	Link             string             `json:"link"`
	WeeklyPost       bool               `json:"weeklyPost"`
}

type WeeklyToolkitPost struct {
	ToolkitTitle       string `json:"toolkitTitle,omitempty"`
	WeeklyToolkitPosts []Post `json:"weeklyToolkitPosts,omitempty"`
}

func (newToolkitPost *Post) SaveToolkitPost(c context.Context) (Post, error) {
	mToolkitPost, err := toolkitPostCollection.InsertOne(c, newToolkitPost)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return Post{}, errors.New("Toolkit Post Title Already Exists")
		}
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"title": 1}, Options: opt}

	if _, err := toolkitPostCollection.Indexes().CreateOne(c, index); err != nil {
		return Post{}, errors.New("could not create index of Toolkit Post Title")
	}

	var curToolkitPost Post
	toolkitPostCollection.FindOne(c, bson.M{"_id": mToolkitPost.InsertedID}).Decode(&curToolkitPost)
	return curToolkitPost, nil
}

func GetAllPostsInToolkit(c context.Context) ([]Post, error) {

	var toolkitPosts []Post

	result, err := toolkitPostCollection.Find(c, bson.M{})
	if err != nil {
		return []Post{}, errors.New(SerendipityResponse.ERROR_GETALL_TOOLKITPOSTS)
	}

	defer result.Close(c)

	for result.Next(c) {
		var singleToolkitPost Post
		if err := result.Decode(&singleToolkitPost); err != nil {
			return []Post{}, errors.New(SerendipityResponse.ERROR_DECODE_TOOLKITPOST)
		}
		toolkitPosts = append(toolkitPosts, singleToolkitPost)
	}
	return toolkitPosts, nil
}

func GetAllToolkitPostsInToolkit(limit int, page int, sortField string, sortOrder string, objId int, c context.Context) ([]Post, error) {

	skipNumber := (page - 1) * limit
	pageOptions := options.Find()
	if sortOrder == "ascend" {
		pageOptions.SetSort(bson.D{primitive.E{Key: sortField, Value: 1}})
	} else if sortOrder == "descend" {
		pageOptions.SetSort(bson.D{primitive.E{Key: sortField, Value: -1}})
	}
	pageOptions.SetSkip(int64(skipNumber))
	pageOptions.SetLimit(int64(limit))

	var toolkitPosts []Post
	result, err := toolkitPostCollection.Find(c, bson.M{"toolkittype": objId}, pageOptions)
	if err != nil {
		return []Post{}, errors.New(SerendipityResponse.ERROR_GETALL_TOOLKITPOSTS)
	}

	defer result.Close(c)

	for result.Next(c) {
		var singleToolkitPost Post
		if err := result.Decode(&singleToolkitPost); err != nil {
			return []Post{}, errors.New(SerendipityResponse.ERROR_DECODE_TOOLKITPOST)
		}
		toolkitPosts = append(toolkitPosts, singleToolkitPost)
	}
	return toolkitPosts, nil

}

func GetToolkitPostOne(objId primitive.ObjectID, c context.Context) (Post, error) {
	var getToolkitPost Post
	err := toolkitPostCollection.FindOne(c, bson.M{"id": objId}).Decode(&getToolkitPost)
	if err != nil {
		return Post{}, errors.New(SerendipityResponse.ERROR_GET_TOOLKITPOST_ONE)
	}
	if getToolkitPost.ID.Hex() == "000000000000000000000000" {
		return Post{}, errors.New(SerendipityResponse.NOT_EXIST_TOOLKITPOST)
	}
	return getToolkitPost, nil
}

func UpdateToolkitPostWithID(objId primitive.ObjectID, update bson.M, c context.Context) (Post, error) {
	result, err := toolkitPostCollection.UpdateOne(c, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return Post{}, errors.New("Toolkit Post Title Already Exists")
		}
	} else if result.MatchedCount < 1 {
		return Post{}, errors.New(SerendipityResponse.WARNING_UPDATE_TOOLKITPOST)
	}
	return Post{}, nil
}

func DeleteToolkitPostOne(objId primitive.ObjectID, c context.Context) (Post, error) {
	result, err := toolkitPostCollection.DeleteOne(c, bson.M{"id": objId})
	if err != nil {
		return Post{}, errors.New(SerendipityResponse.DELETE_POST_ERROR)
	}
	if result.DeletedCount < 1 {
		return Post{}, errors.New(SerendipityResponse.NOT_FOUND_POST_DELETE)
	}
	return Post{}, nil
}

func GetAllWeeklyPostsInToolkit(c context.Context) ([]WeeklyToolkitPost, error) {
	var weeklyToolkitPosts []WeeklyToolkitPost

	nutritionResult, err := toolkitPostCollection.Find(c, bson.M{"toolkittype": 1, "weeklypost": true})
	if err != nil {
		return []WeeklyToolkitPost{}, errors.New(SerendipityResponse.ERROR_GET_WEEKLYTOOLKITPOSTS)
	}
	healthyResult, err := toolkitPostCollection.Find(c, bson.M{"toolkittype": 2, "weeklypost": true})
	if err != nil {
		return []WeeklyToolkitPost{}, errors.New(SerendipityResponse.ERROR_GET_WEEKLYTOOLKITPOSTS)
	}
	recipeResult, err := toolkitPostCollection.Find(c, bson.M{"toolkittype": 3, "weeklypost": true})
	if err != nil {
		return []WeeklyToolkitPost{}, errors.New(SerendipityResponse.ERROR_GET_WEEKLYTOOLKITPOSTS)
	}
	movementResult, err := toolkitPostCollection.Find(c, bson.M{"toolkittype": 4, "weeklypost": true})
	if err != nil {
		return []WeeklyToolkitPost{}, errors.New(SerendipityResponse.ERROR_GET_WEEKLYTOOLKITPOSTS)
	}
	meditationResult, err := toolkitPostCollection.Find(c, bson.M{"toolkittype": 5, "weeklypost": true})
	if err != nil {
		return []WeeklyToolkitPost{}, errors.New(SerendipityResponse.ERROR_GET_WEEKLYTOOLKITPOSTS)
	}
	defer nutritionResult.Close(c)
	defer healthyResult.Close(c)
	defer recipeResult.Close(c)
	defer movementResult.Close(c)
	defer meditationResult.Close(c)

	var singleWeeklyToolkitPostForNutrition WeeklyToolkitPost
	for nutritionResult.Next(c) {
		var singleToolkitPost Post
		singleWeeklyToolkitPostForNutrition.ToolkitTitle = "Nutrition Resources"
		if err := nutritionResult.Decode(&singleToolkitPost); err != nil {
			return []WeeklyToolkitPost{}, errors.New(SerendipityResponse.ERROR_DECODE_TOOLKITPOST)
		}
		singleWeeklyToolkitPostForNutrition.WeeklyToolkitPosts = append(singleWeeklyToolkitPostForNutrition.WeeklyToolkitPosts, singleToolkitPost)
	}
	weeklyToolkitPosts = append(weeklyToolkitPosts, singleWeeklyToolkitPostForNutrition)
	var singleWeeklyToolkitPostForHealthy WeeklyToolkitPost
	for healthyResult.Next(c) {
		var singleToolkitPost Post
		singleWeeklyToolkitPostForHealthy.ToolkitTitle = "Healthy Home"
		if err := healthyResult.Decode(&singleToolkitPost); err != nil {
			return []WeeklyToolkitPost{}, errors.New(SerendipityResponse.ERROR_DECODE_TOOLKITPOST)
		}
		singleWeeklyToolkitPostForHealthy.WeeklyToolkitPosts = append(singleWeeklyToolkitPostForHealthy.WeeklyToolkitPosts, singleToolkitPost)
	}
	weeklyToolkitPosts = append(weeklyToolkitPosts, singleWeeklyToolkitPostForHealthy)
	var singleWeeklyToolkitPostForRecipe WeeklyToolkitPost
	for recipeResult.Next(c) {
		var singleToolkitPost Post
		singleWeeklyToolkitPostForRecipe.ToolkitTitle = "Recipe Archive"
		if err := recipeResult.Decode(&singleToolkitPost); err != nil {
			return []WeeklyToolkitPost{}, errors.New(SerendipityResponse.ERROR_DECODE_TOOLKITPOST)
		}
		singleWeeklyToolkitPostForRecipe.WeeklyToolkitPosts = append(singleWeeklyToolkitPostForRecipe.WeeklyToolkitPosts, singleToolkitPost)
	}
	weeklyToolkitPosts = append(weeklyToolkitPosts, singleWeeklyToolkitPostForRecipe)
	var singleWeeklyToolkitPostForMovement WeeklyToolkitPost
	for movementResult.Next(c) {
		var singleToolkitPost Post
		singleWeeklyToolkitPostForMovement.ToolkitTitle = "Movement Archive"
		if err := movementResult.Decode(&singleToolkitPost); err != nil {
			return []WeeklyToolkitPost{}, errors.New(SerendipityResponse.ERROR_DECODE_TOOLKITPOST)
		}
		singleWeeklyToolkitPostForMovement.WeeklyToolkitPosts = append(singleWeeklyToolkitPostForMovement.WeeklyToolkitPosts, singleToolkitPost)
	}
	weeklyToolkitPosts = append(weeklyToolkitPosts, singleWeeklyToolkitPostForMovement)
	var singleWeeklyToolkitPostForMeditation WeeklyToolkitPost
	for meditationResult.Next(c) {
		var singleToolkitPost Post
		singleWeeklyToolkitPostForMeditation.ToolkitTitle = "Meditation Archive"
		if err := meditationResult.Decode(&singleToolkitPost); err != nil {
			return []WeeklyToolkitPost{}, errors.New(SerendipityResponse.ERROR_DECODE_TOOLKITPOST)
		}
		singleWeeklyToolkitPostForMeditation.WeeklyToolkitPosts = append(singleWeeklyToolkitPostForMeditation.WeeklyToolkitPosts, singleToolkitPost)
	}
	weeklyToolkitPosts = append(weeklyToolkitPosts, singleWeeklyToolkitPostForMeditation)
	return weeklyToolkitPosts, nil
}
