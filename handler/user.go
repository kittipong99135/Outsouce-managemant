package handler

import (
	service "osm/service/user_service"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	// userSrv service.UserService
	userSrv service.UserService
}

func NewUserHandler(userSrv service.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) GetAll(c *fiber.Ctx) error {
	result, err := h.userSrv.SrvGetAllUser(c)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"result":       "Fail",
			"errorMessage": err.Error(),
			"data":         nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result":       "Success",
		"errorMessage": nil,
		"data":         result,
	})

}

func (h userHandler) GetById(c *fiber.Ctx) error {

	result, err := h.userSrv.SrvGetById(c, c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"result":       "Fail",
			"errorMessage": err.Error(),
			"data":         nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result":       "Success",
		"errorMessage": nil,
		"data":         result,
	})
}
