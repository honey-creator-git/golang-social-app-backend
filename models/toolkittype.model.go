package models

import (
	"context"
	"errors"
	"serendipity_backend/SerendipityResponse"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToolkitType struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	Title            string             `json:"title,omitempty"`            // -> Nutrition Resources, Foundation Resources, Recipe Resources, Movement Archive, Meditation Archive
	CoverLetterImage string             `json:"coverletterimage,omitempty"` //-> image url
	SortType         []string           `json:"sortType"`
	Type             int                `json:"type,omitempty"` //-> 1: Nutrition Resources, 2: Foundation Resources, 3: Recipe Archive, 4: Movement Archive, 5: Meditation Archive
}

func (newToolkitType *ToolkitType) SaveToolkit(c context.Context) (ToolkitType, error) {
	mToolkit, err := toolkitTypeCollection.InsertOne(c, newToolkitType)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return ToolkitType{}, errors.New("Toolkit Type Already Exists")
		}
	}

	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"type": 1}, Options: opt}

	if _, err := toolkitTypeCollection.Indexes().CreateOne(c, index); err != nil {
		return ToolkitType{}, errors.New("could not create index of toolkit type")
	}

	var curToolkit ToolkitType
	toolkitTypeCollection.FindOne(c, bson.M{"_id": mToolkit.InsertedID}).Decode(&curToolkit)
	return curToolkit, nil
}

func GetToolkitTypeWithID(objId primitive.ObjectID, c context.Context) (ToolkitType, error) {
	var toolkitType ToolkitType
	err := toolkitTypeCollection.FindOne(c, bson.M{"id": objId}).Decode(&toolkitType)
	if err != nil {
		return ToolkitType{}, errors.New(SerendipityResponse.ERROR_GET_TOOLKITTYPE)
	}
	return toolkitType, nil
}

func UpdateToolkitType(objId primitive.ObjectID, update bson.M, c context.Context) (ToolkitType, error) {
	result, err := toolkitTypeCollection.UpdateOne(c, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return ToolkitType{}, errors.New(SerendipityResponse.ERROR_UPDATE_TOOLKITTYPE)
	}
	if result.MatchedCount < 1 {
		return ToolkitType{}, errors.New(SerendipityResponse.WARNING_UPDATE_TOOLKITTYPE)
	}
	return ToolkitType{}, nil
}

// func FindAllToolkitTypes(limit int, page int, sortField string, sortOrder string, c context.Context) ([]ToolkitType, error) {
func FindAllToolkitTypes(c context.Context) ([]ToolkitType, error) {
	// skipNumber := (page - 1) * limit
	// pageOptions := options.Find()
	// if sortOrder == "ascend" {
	// 	pageOptions.SetSort(bson.D{primitive.E{Key: sortField, Value: 1}})
	// } else if sortOrder == "descend" {
	// 	pageOptions.SetSort(bson.D{primitive.E{Key: sortField, Value: -1}})
	// }
	// pageOptions.SetSkip(int64(skipNumber))
	// pageOptions.SetLimit(int64(limit))

	var toolkitTypes []ToolkitType
	// results, err := toolkitTypeCollection.Find(c, bson.M{}, pageOptions)
	results, err := toolkitTypeCollection.Find(c, bson.M{})
	if err != nil {
		return []ToolkitType{}, errors.New(SerendipityResponse.ERROR_GET_TOOLKITTYPES)
	}

	defer results.Close(c)

	for results.Next(c) {
		var singleToolkitType ToolkitType
		if err := results.Decode(&singleToolkitType); err != nil {
			return []ToolkitType{}, errors.New(SerendipityResponse.ERROR_DECODE_TOOLKITTYPE)
		}
		toolkitTypes = append(toolkitTypes, singleToolkitType)
	}
	return toolkitTypes, nil
}

func DeleteToolkitOne(objId int, c context.Context) (ToolkitType, error) {
	result, err := toolkitTypeCollection.DeleteOne(c, bson.M{"type": objId})
	if err != nil {
		return ToolkitType{}, errors.New(SerendipityResponse.DELETE_TOOLKIT_ERROR)
	}
	if result.DeletedCount < 1 {
		return ToolkitType{}, errors.New(SerendipityResponse.NOT_FOUND_TOOLKIT_TYPE)
	}
	resultt, errr := toolkitPostCollection.DeleteMany(c, bson.M{"toolkittype": objId})
	if errr != nil {
		return ToolkitType{}, errors.New(SerendipityResponse.DELETE_TOOLKIT_POSTS)
	}
	if resultt.DeletedCount < 1 {
		return ToolkitType{}, nil
	}
	return ToolkitType{}, nil
}
