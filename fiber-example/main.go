package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	app.Get("/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
