package database

import (
	"context"
	"osm/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// "context"
// "gomon/models"

const dbName = "SODdb"
const mongoURL = "mongodb://127.0.0.1:27017/" + dbName

func MongoInit() (*models.MongoInstance, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 30*time.Second)
	_ = cancle
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)
	if err != nil {
		return nil, err
	}

	mg := models.MongoInstance{
		Client: client,
		Ctx:    ctx,
		Db:     db,
	}

	return &mg, nil
}
