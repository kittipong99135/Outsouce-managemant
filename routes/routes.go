package routes

import (
	"osm/handler"
	"osm/models"

	auth_repository "osm/repository/auth_repo"
	user_repository "osm/repository/user_repo"

	auth_service "osm/service/auth_service"
	user_service "osm/service/user_service"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, mgConn *models.MongoInstance) {

	api := app.Group("/api")
	v1 := api.Group("/v1")

	authRep := auth_repository.NewAuthRepository()
	authSrv := auth_service.NewAuthRepository(authRep)
	authHan := handler.NewAuthService(authSrv)

	auth := v1.Group("/auth")

	auth.Post("/login", authHan.Login)

	userRep := user_repository.NewUserRepository(mgConn)
	userSrv := user_service.NewUserService(userRep)
	userHan := handler.NewUserHandler(userSrv)

	user := v1.Group("/user")

	user.Get("/getall", userHan.GetAll)
	user.Get("/getall/:id", userHan.GetById)

}
