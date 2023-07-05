package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})

	// Parameter
	// GET http://localhost:8080/hi/fian
	app.Get("/hi/:name", func(c *fiber.Ctx) error {
		return c.SendString("Hello " + c.Params("name"))
		// => Hello fian
	})

	// Parameter Optional
	// GET http://localhost:8080/hai/ogut
	app.Get("/hai/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hai " + c.Params("name"))
		}
		// => Hai ogut
		return c.SendString("Where is user?")
	})

	// Wildcard
	// GET http://localhost:8080/fian/smart
	app.Get("/:name/*", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name") + " " + c.Params("*"))
		}
		// => Hello fian smart
		return c.SendString("Whre is userr?")
	})

	app.Listen(":8080")
}
