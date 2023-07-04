package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Middleware mengangani error secara global
	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				// Tangan error dan return response
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": fmt.Sprintf("%v", r),
				})
			}
		}()
		return c.Next()
	})

	// Error Handling di Handler
	app.Get("/err", func(c *fiber.Ctx) error {
		// Contoh error
		err := fmt.Errorf("Something went wrong!")

		// Caught error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("%v", err),
			})
		}
		return c.SendString("No error")
	})
	app.Listen(":8080")
}
