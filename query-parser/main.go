package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Name     string   `json:"name"`
	Pass     string   `json:"pass"`
	Products []string `json:"products"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		p := new(Person)

		if err := c.QueryParser(p); err != nil {
			return err
		}

		log.Println(p.Name)
		log.Println(p.Pass)
		log.Println(p.Products)
		return nil
	})
	app.Listen(":8080")
}
