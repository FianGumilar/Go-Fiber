package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Menerapkan beberapa stack
	app.Get("/api/:name", func(c *fiber.Ctx) error {
		return c.SendString("Access API " + c.Params("name"))
	})

	app.Post("/api", func(c *fiber.Ctx) error {
		return c.SendString("Create new API")
	})

	stack := app.Stack()

	data, _ := json.MarshalIndent(stack, "", " ")
	fmt.Println(string(data))

	app.Listen(":8080")

}
