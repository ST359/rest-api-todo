package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/tasks", func(c *fiber.Ctx) error {
		return c.SendString("tasks")
	})
	app.Post("/tasks", func(c *fiber.Ctx) error {
		return c.SendString("post task")
	})
	app.Put("/tasks/:id", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("id"))
	})
	app.Delete("/tasks/:id", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("id"))
	})
	app.Listen(":3000")
}
