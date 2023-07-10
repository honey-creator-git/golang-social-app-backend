package models

import (
	"serendipity_backend/configs"

	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var toolkitTypeCollection *mongo.Collection = configs.GetCollection(configs.DB, "toolkit")
var toolkitPostCollection *mongo.Collection = configs.GetCollection(configs.DB, "toolkitPost")
var forumTypeCollection *mongo.Collection = configs.GetCollection(configs.DB, "forumType")
var forumPostCollection *mongo.Collection = configs.GetCollection(configs.DB, "forumPost")
var foundationCollection *mongo.Collection = configs.GetCollection(configs.DB, "foundations")
var marketplaceCollection *mongo.Collection = configs.GetCollection(configs.DB, "marketplace")
var marketplaceItemCollection *mongo.Collection = configs.GetCollection(configs.DB, "marketplaceItem")

// var forumPostCommentCollection *mongo.Collection = configs.GetCollection(configs.DB, "forumPostComment")
