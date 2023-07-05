package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CutomConfig() {
	app := fiber.New()

	app.Static("/", "./public", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})
}

func main() {

	//Custom Config
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "TestApp v1.0.0",
	})
	fmt.Println(app)

	if fiber.IsChild() {
		fmt.Println("I'm in parent process")
	} else {
		fmt.Println("I'm in child process")
	}
}
