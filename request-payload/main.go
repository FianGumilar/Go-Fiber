package main

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=0,lte=80"`
}

func main() {
	app := fiber.New()

	// Konfigurasi validator
	validate := validator.New()

	app.Post("/user", func(c *fiber.Ctx) error {
		user := new(User)

		// binding data dari body request
		if err := c.BodyParser(user); err != nil {
			return err
		}

		// validasi payload
		if err := validate.Struct(user); err != nil {
			//Jika terdapat error, kembalikan dengan response error
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		log.Printf("User %+v", user)

		return c.JSON(user)
	})
	log.Fatal(app.Listen(":8080"))
}
