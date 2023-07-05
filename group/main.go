package main

import "github.com/gofiber/fiber/v2"

// fitur Group untuk mengelompokan rute-rute dan menerapkan middleware untuk seluruh grup

func main() {
	app := fiber.New()

	// Menerapkan middleware untuk seluruh grup
	app.Use("/api", func(c *fiber.Ctx) error {
		// Middleware golabal
		return c.Next()
	})

	// Membuat group
	api := app.Group("/api", func(c *fiber.Ctx) error {
		return c.Next()
	})

	api.Get("/users", func(c *fiber.Ctx) error {
		// Handler untuk /api/users
		return c.SendString("API Users")
	})

	api.Get("/posts", func(c *fiber.Ctx) error {
		// Handler untuk /api/posts
		return c.SendString("API Posts")
	})

	app.Listen(":8080")
}
