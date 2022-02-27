package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		// https://docs.gofiber.io/api/fiber
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Fiber v2 App v1.0",
	})

	if !fiber.IsChild() {
		fmt.Println("I'm the parent process")
	} else {
		fmt.Println("I'm a child process")
	}

	app.Static("/static", "./public")

	// app.Get("/:value", func(c *fiber.Ctx) error {
	// 	return c.SendString("value: " + c.Params("value"))
	// })

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(782, "Custom error message")
	})

	app.Listen(":3000")
}
