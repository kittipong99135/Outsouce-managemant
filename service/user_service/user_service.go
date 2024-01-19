package service

import (
	"errors"
	"osm/models"
	repository "osm/repository/user_repo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userService struct {
	// userRepo repository.UserRepository
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{userRepo: userRepo}
}

func (s userService) SrvGetAllUser(c *fiber.Ctx) ([]models.UserResponse, error) {
	query_result, err := s.userRepo.GetAll()
	if err != nil {
		return nil, errors.New("not convert data to struct")
	}

	var user []models.UserResponse
	if err := query_result.All(c.Context(), &user); err != nil {
		return nil, errors.New("not convert from employee")
	}
	return user, nil
}

func (s userService) SrvGetById(c *fiber.Ctx, params string) ([]models.UserResponse, error) {
	_id, err := primitive.ObjectIDFromHex(params)
	if err != nil {
		return nil, errors.New("missing params number")
	}

	query_result, err := s.userRepo.GetById(_id)
	if err != nil {
		return nil, errors.New("id not found")
	}

	var user []models.UserResponse

	if err := query_result.All(c.Context(), &user); err != nil {
		return nil, errors.New("not convert from employee")
	}

	return user, nil
}
