package repository

import (
	"osm/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StaffRepository interface {
	Create(*fiber.Ctx, *models.Staff) (*mongo.InsertOneResult, error)
	GetAll(*fiber.Ctx) (*mongo.Cursor, error)
	GetById(*fiber.Ctx, primitive.ObjectID) (*mongo.Cursor, error)
	Update(*fiber.Ctx, primitive.ObjectID) error
	Remove(*fiber.Ctx, primitive.ObjectID) error
}
