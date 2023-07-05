package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Cookie struct {
	Name        string    `json:"name"`
	Value       string    `json:"value"`
	Path        string    `json:"path"`
	Domain      string    `json:"domain"`
	MaxAge      int       `json:"max_age`
	Expires     time.Time `json:"expires"`
	Secure      bool      `json:"secure"`
	HTTPOnly    bool      `json:"http_only"`
	SameSite    string    `json:"same_site"`
	SessionOnly bool      `json:"session_only"`
}

func main() {
	app := fiber.New()

	app.Get("/cookie", func(c *fiber.Ctx) error {
		cookie := new(fiber.Cookie)
		cookie.Name = "cookieData"
		cookie.Value = "12345678"
		cookie.Expires = time.Now().Add(5 * time.Minute)

		c.Cookie(cookie)
		return nil
	})

	app.Listen(":8080")

}
