package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetAll() (*mongo.Cursor, error)
	GetById(primitive.ObjectID) (*mongo.Cursor, error)
}
