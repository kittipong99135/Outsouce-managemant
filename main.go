package main

import (
	"fmt"
	"os/exec"
	"osm/database"
	"osm/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	mgConn, err := database.MongoInit()
	if err != nil {
		fmt.Println("Can't connetc MongoDB.")
	}

	fmt.Println("Success to connected MongoDB.")

	app := fiber.New()

	routes.Routes(app, mgConn)

	app.Get("/", func(c *fiber.Ctx) error {

		redirectURL := "https://one.th/oauth/login"
		openBrowser(redirectURL)

		return nil
	})

	app.Listen(":3000")
}

func openBrowser(url string) {
	cmd := exec.Command("cmd", "/c", "start", url)
	cmd.Start()
}
