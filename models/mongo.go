package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoInstance struct {
	Client *mongo.Client
	Ctx    context.Context
	Db     *mongo.Database
}
