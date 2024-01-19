package repository

import (
	"osm/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	mgConn *models.MongoInstance
}

func NewUserRepository(mgConn *models.MongoInstance) UserRepository {
	return userRepository{mgConn: mgConn}
}

func (r userRepository) GetAll() (*mongo.Cursor, error) {
	collection := r.mgConn.Db.Collection("users")

	query := bson.D{{}}
	query_result, err := collection.Find(r.mgConn.Ctx, query)
	if err != nil {
		return nil, err
	}

	return query_result, nil
}

func (r userRepository) GetById(_id primitive.ObjectID) (*mongo.Cursor, error) {
	collection := r.mgConn.Db.Collection("users")

	query := bson.D{{Key: "_id", Value: _id}}
	query_result, err := collection.Find(r.mgConn.Ctx, query)
	if err != nil {
		return nil, err
	}

	return query_result, nil
}
