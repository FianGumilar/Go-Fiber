package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int32  `json:"age" validate:"gte=0,lte=80"`
}

func main() {
	app := fiber.New()

	// konfigurasi validator
	validate := validator.New()

	app.Post("/user", func(c *fiber.Ctx) error {
		user := new(User)

		// binding dari body request
		if err := c.BodyParser(user); err != nil {
			return err
		}

		// Validasi payload
		if err := validate.Struct(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		log.Printf("User %+v", user)

		return c.JSON(user)
	})

	log.Fatal(app.Listen(":8080"))
}
