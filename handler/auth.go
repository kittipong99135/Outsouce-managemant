package handler

import (
	"osm/models"
	service "osm/service/auth_service"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authSrv service.AuthService
}

func NewAuthService(authSrv service.AuthService) authHandler {
	return authHandler{authSrv: authSrv}
}

func (h authHandler) Login(c *fiber.Ctx) error {
	var login_body models.Login_body
	if err := c.BodyParser(&login_body); err != nil {
		return c.Status(400).SendString("can't compile bodyparser")
	}

	result, err := h.authSrv.SrvLogin(&login_body)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"result":       "Fail",
			"data":         nil,
			"errorMessage": err.Error(),
			"code":         400,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"result":       "Success",
		"data":         result,
		"errorMessage": nil,
		"code":         200,
	})
}
