package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/:name", func(c *fiber.Ctx) error {
		return c.SendString("Access API " + c.Params("name"))
	}).Name("Get Name")

	app.Post("/api", func(c *fiber.Ctx) error {
		return c.SendString("Create new API")
	}).Name("Post Api")

	stack := app.Stack()
	data, _ := json.MarshalIndent(stack, "", " ")
	fmt.Print(string(data))

	app.Listen(":8080")
}
