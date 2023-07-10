package models

import (
	"context"
	"errors"
	"fmt"
	"serendipity_backend/SerendipityResponse"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type ForumPostCreator struct {
// 	ID primitive.ObjectID `json:"id,omitempty"`
// 	FirstName string `json:"firstName,omitempty"`
// 	LastName string `json:"lastName,omitempty"`
// 	Avatar string `json:"avatar"`
// 	Follows []primitive
// }
type Memeber struct {
	ID string `json:"id,omitempty"`
}

type ForumPostEmotion struct {
	Like    []Memeber `json:"like,omitempty"`
	Dislike []Memeber `json:"dislike,omitempty"`
}

type Emotion struct {
	Like    int `json:"like,omitempty"`
	Dislike int `json:"dislike,omitempty"`
}

type ForumType struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	Title            string             `json:"title,omitempty"`
	CoverLetterImage string             `json:"coverLetterImage,omitempty"`
	ForumType        int                `json:"forumType,omitempty"`
}

type Comment struct {
	ID          primitive.ObjectID `json:"id,omitempty"`
	Description string             `json:"description,omitempty"`
	PostedAt    string             `json:"postedAt,omitempty"`
	PostedBy    primitive.ObjectID `json:"postedBy,omitempty"`
	PostId      primitive.ObjectID `json:"postId,omitEmpty"`
	Emotions    Emotion            `json:"emotions,omitempty"`
}

type ForumPost struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	Title            string             `json:"title,omitempty"`
	CoverLetterImage string             `json:"coverLetterImage,omitempty"`
	Description      string             `json:"description,omitemtpy"`
	CreatedAt        string             `json:"createdAt"`
	CreatedBy        primitive.ObjectID `json:"createdBy,omitempty"`
	Comments         []Comment          `json:"comments"`
	VisitCount       int                `json:"visitCount"`
	Emotions         ForumPostEmotion   `json:"emotions"`
	ForumType        int                `json:"forumType,omitempty"`
}

func (newForumType *ForumType) SaveForumType(c context.Context) (ForumType, error) {
	mForumType, err := forumTypeCollection.InsertOne(c, newForumType)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return ForumType{}, errors.New("Forum Type alrady exists")
		}
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"title": 1}, Options: opt}

	if _, err := forumTypeCollection.Indexes().CreateOne(c, index); err != nil {
		return ForumType{}, errors.New("could not create index of forum type title")
	}

	var currForumType ForumType
	forumTypeCollection.FindOne(c, bson.M{"_id": mForumType.InsertedID}).Decode(&currForumType)
	return currForumType, nil
}

func (newForumPost *ForumPost) SaveForumPost(c context.Context) (ForumPost, error) {
	mForumPost, err := forumPostCollection.InsertOne(c, newForumPost)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return ForumPost{}, errors.New("This Forum Post alrady exists")
		}
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"title": 1}, Options: opt}

	if _, err := forumPostCollection.Indexes().CreateOne(c, index); err != nil {
		return ForumPost{}, errors.New("could not create index of forum post title")
	}

	var currForumPost ForumPost
	forumPostCollection.FindOne(c, bson.M{"_id": mForumPost.InsertedID}).Decode(&currForumPost)
	return currForumPost, nil
}

func AddCommentForForumPost(objPostId primitive.ObjectID, addedComment bson.M, c context.Context) ([]Comment, error) {
	fmt.Printf("Param for add new comment => %v", addedComment)
	result, err := forumPostCollection.UpdateOne(c, bson.M{"id": objPostId}, bson.M{"$set": addedComment})
	if err != nil {
		return []Comment{}, errors.New(SerendipityResponse.ERROR_COMMENT_POST_FORUM)
	}
	if result.MatchedCount == 1 {
		var updatedForumPost ForumPost
		err := forumPostCollection.FindOne(c, bson.M{"id": objPostId}).Decode(&updatedForumPost)
		if err != nil {
			return []Comment{}, errors.New(SerendipityResponse.ERROR_GET_UPDATED_FORUM_POST)
		}
		fmt.Printf("Updated Forum Post Comments => %v", updatedForumPost.Comments)
		return updatedForumPost.Comments, nil
	} else {
		return []Comment{}, errors.New(SerendipityResponse.ERROR_COMMENT_ADD_FORUM_POST)
	}
}

func GetForumPostWithID(objId primitive.ObjectID, c context.Context) (ForumPost, error) {
	var forumPost ForumPost
	err := forumPostCollection.FindOne(c, bson.M{"id": objId}).Decode(&forumPost)
	if err != nil {
		return ForumPost{}, errors.New(SerendipityResponse.ERROR_GET_POST_FORUM_ID)
	}
	if forumPost.ID.Hex() == "000000000000000000000000" {
		return ForumPost{}, errors.New(SerendipityResponse.NOT_EXIST_FORUMPOST)
	}
	return forumPost, nil
}

// func GetToolkitPostOne(objId primitive.ObjectID, c context.Context) (Post, error) {
// 	var getToolkitPost Post
// 	err := toolkitPostCollection.FindOne(c, bson.M{"id": objId}).Decode(&getToolkitPost)
// 	if err != nil {
// 		return Post{}, errors.New(SerendipityResponse.ERROR_GET_TOOLKITPOST_ONE)
// 	}
// 	if getToolkitPost.ID.Hex() == "000000000000000000000000" {
// 		return Post{}, errors.New(SerendipityResponse.NOT_EXIST_TOOLKITPOST)
// 	}
// 	return getToolkitPost, nil
// }

func UpdateForumPostComments(objId primitive.ObjectID, update bson.M, c context.Context) (ForumPost, error) {
	result, err := forumPostCollection.UpdateOne(c, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return ForumPost{}, errors.New(SerendipityResponse.ERROR_UPDATE_POST_COMMENTS_FORUM_ID)
	}
	if result.MatchedCount < 1 {
		return ForumPost{}, errors.New(SerendipityResponse.WARNING_UPDATE_POST_COMMENTS_FORUM_ID)
	}
	return ForumPost{}, nil
}

func UpdateForumPostWithPostIDUserID(objPostId primitive.ObjectID, userId primitive.ObjectID, update bson.M, c context.Context) (ForumPost, error) {
	result, err := forumPostCollection.UpdateOne(c, bson.M{"id": objPostId, "createdby": userId}, bson.M{"$set": update})
	if err != nil {
		return ForumPost{}, errors.New(SerendipityResponse.ERROR_UPDATE_FORUM_POST)
	}
	if result.MatchedCount < 1 {
		return ForumPost{}, errors.New(SerendipityResponse.WARNING_UPDATE_FORUM_POST)
	}
	return ForumPost{}, nil
}

func UpdateForumPostWithPostID(objPostId primitive.ObjectID, update bson.M, c context.Context) (ForumPost, error) {
	result, err := forumPostCollection.UpdateOne(c, bson.M{"id": objPostId}, bson.M{"$set": update})
	if err != nil {
		return ForumPost{}, errors.New(SerendipityResponse.ERROR_UPDATE_FORUM_POST)
	}
	if result.MatchedCount < 1 {
		return ForumPost{}, errors.New(SerendipityResponse.WARNING_UPDATE_FORUM_POST)
	}
	return ForumPost{}, nil
}
func GetAllForumPostsWithForumID(c context.Context) ([]ForumPost, error) {
	var forumPosts []ForumPost

	result, err := forumPostCollection.Find(c, bson.M{})
	if err != nil {
		return []ForumPost{}, errors.New(SerendipityResponse.ERROR_GET_ALL_FORUM_POSTS)
	}

	defer result.Close(c)

	for result.Next(c) {
		var SingleForumPost ForumPost
		if err := result.Decode(&SingleForumPost); err != nil {
			return []ForumPost{}, errors.New(SerendipityResponse.ERROR_DECODE_FORUM_POST)
		}
		forumPosts = append(forumPosts, SingleForumPost)
	}

	return forumPosts, nil
}

func GetForumTypes(c context.Context) ([]ForumType, error) {
	var forumTypes []ForumType
	results, err := forumTypeCollection.Find(c, bson.M{})
	if err != nil {
		return []ForumType{}, errors.New(SerendipityResponse.ERROR_GET_FORUMTYPES)
	}

	defer results.Close(c)

	for results.Next(c) {
		var SingleForumType ForumType
		if err := results.Decode(&SingleForumType); err != nil {
			return []ForumType{}, errors.New(SerendipityResponse.ERROR_DECODE_FORUMTYPE)
		}
		forumTypes = append(forumTypes, SingleForumType)
	}

	return forumTypes, nil
}

func GetForumPostWithPostID(objId primitive.ObjectID, c context.Context) (ForumPost, error) {
	var SingleForumPost ForumPost
	err := forumPostCollection.FindOne(c, bson.M{"id": objId}).Decode(&SingleForumPost)
	if err != nil {
		return ForumPost{}, errors.New(SerendipityResponse.ERROR_GET_FORUMPOST_WITHID)
	}
	return SingleForumPost, nil
}

func DeleteForumPostOne(objId primitive.ObjectID, c context.Context) error {
	result, err := forumPostCollection.DeleteOne(c, bson.M{"id": objId})

	if err != nil {
		return errors.New(SerendipityResponse.ERROR_DELETE_FORUMPOST)
	}

	if result.DeletedCount < 1 {
		return errors.New(SerendipityResponse.WARNING_DELETE_FORUMPOST)
	}

	return nil
}

func DeleteForumOne(objId int, c context.Context) (ForumType, error) {
	result, err := forumTypeCollection.DeleteOne(c, bson.M{"forumtype": objId})
	if err != nil {
		return ForumType{}, errors.New(SerendipityResponse.ERROR_DELETE_FORUM_TYPE)
	}
	if result.DeletedCount < 1 {
		return ForumType{}, errors.New(SerendipityResponse.NOT_FOUND_FORUM_TYPE)
	}
	resultt, errr := forumPostCollection.DeleteMany(c, bson.M{"forumtype": objId})
	if errr != nil {
		return ForumType{}, errors.New(SerendipityResponse.DELETE_FORUM_POSTS)
	}
	if resultt.DeletedCount < 1 {
		return ForumType{}, nil
	}
	return ForumType{}, nil
}

func FindForumTypeWithID(objId primitive.ObjectID, c context.Context) (ForumType, error) {
	var forumType ForumType
	err := forumTypeCollection.FindOne(c, bson.M{"id": objId}).Decode(&forumType)

	if err != nil {
		return ForumType{}, errors.New(SerendipityResponse.ERROR_GET_FORUM_TYPE)
	}

	return forumType, nil
}

func UpdateForumType(objId primitive.ObjectID, update bson.M, c context.Context) (ForumType, error) {
	result, err := forumTypeCollection.UpdateOne(c, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		return ForumType{}, errors.New(SerendipityResponse.ERROR_FORUM_UPDATE)
	}

	if result.MatchedCount < 1 {
		return ForumType{}, errors.New(SerendipityResponse.FORUM_NOT_EXISTED)
	}

	return ForumType{}, nil
}
