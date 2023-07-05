package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Parsing URL Parameters
	app.Get("/user/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		data := fmt.Sprintf("Hello %s", name)

		return c.SendString(data)
	})

	// Wildcard
	app.Get("/user/:name/*", func(c *fiber.Ctx) error {
		name := c.Params("name")
		message := c.Params("*")

		data := fmt.Sprintf("Hello %s, I have message from you %s", name, message)

		return c.SendString(data)
	})

	app.Listen(":8080")
}
