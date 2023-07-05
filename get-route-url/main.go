package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define route with parameters
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		userID := c.Params("id")
		url := c.URI("user", userID) // mengahsilkan rute "URL" dengan parameter id
		return c.SendString(fmt.Sprintf("URL : %s", url))
	}).Name("user")

	app.Listen(":8080")
}
