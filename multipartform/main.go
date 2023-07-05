package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	// Endpoint untuk upload file
	app.Post("/upload", func(c *fiber.Ctx) error {
		// mengambil file dari form
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Bad request")
		}

		// simpan file yang diunngah
		err = c.SaveFile(file, "./uploads"+file.Filename)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		}
		return c.Status(fiber.StatusOK).SendString("successfully uploaded file")
	})
	app.Listen(":8080")
}
