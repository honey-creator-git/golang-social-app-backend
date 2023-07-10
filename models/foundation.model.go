package models

import (
	"context"
	"errors"
	"serendipity_backend/SerendipityResponse"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Foundation struct {
	ID   primitive.ObjectID `json:"id,omitempty"`
	Name string             `json:"name,omitempty"`
}

func GetFoundationWithID(name string, c context.Context) (Foundation, error) {
	var foundation Foundation
	foundationCollection.FindOne(c, bson.M{"name": name}).Decode(&foundation)

	if foundation.ID.Hex() == "000000000000000000000000" {
		return Foundation{}, errors.New(SerendipityResponse.NOT_EXIST_FOUNDATION)
	}

	return foundation, nil
}

func (newFoundation *Foundation) SaveFoundation(c context.Context) (Foundation, error) {
	mFoundation, err := foundationCollection.InsertOne(c, newFoundation)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return Foundation{}, errors.New("Foundation Already Exists")
		}
	}

	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"name": 1}, Options: opt}

	if _, err := foundationCollection.Indexes().CreateOne(c, index); err != nil {
		return Foundation{}, errors.New("Could not create indes of name")
	}

	var curFoundation Foundation
	foundationCollection.FindOne(c, bson.M{"_id": mFoundation.InsertedID}).Decode(&curFoundation)
	return curFoundation, nil
}

func GetAllFoundations(c context.Context) ([]Foundation, error) {
	var foundations []Foundation
	results, err := foundationCollection.Find(c, bson.M{})
	if err != nil {
		return []Foundation{}, errors.New(SerendipityResponse.ERROR_GET_FOUNDATIONS)
	}
	defer results.Close(c)

	for results.Next(c) {
		var singleFoundation Foundation
		if err := results.Decode(&singleFoundation); err != nil {
			return []Foundation{}, errors.New(SerendipityResponse.ERROR_DECODE_FOUNDATION)
		}
		foundations = append(foundations, singleFoundation)
	}
	return foundations, nil
}
