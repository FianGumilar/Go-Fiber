package file_server

import (
	"github.com/gofiber/fiber/v2"
)

func Static() {
	app := fiber.New()

	// Static file server
	app.Static("/", "./files")
	// => http://localhost:8080/test.pdf

	app.Listen(":8080")
}
