package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	// Parameter
	app.Get("/hi/:name", func(c *fiber.Ctx) error {
		return c.SendString("Hello " + c.Params("name"))
	})

	// Parameter Optional
	app.Get("/hai/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hai " + c.Params("name"))
		}
		return c.SendString("Where is user?")
	})

	// Wildcard
	app.Get("/:name/*", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name") + c.Params("*"))
		}
		return c.SendString("Whre is userr?")
	})

	app.Listen(":8080")
}
