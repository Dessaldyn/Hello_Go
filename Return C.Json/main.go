package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	app := fiber.New()

	app.Get("/user", func(c *fiber.Ctx) error {
		user := User{ID: 1, Name: "Muha", Age: 21}
		return c.JSON(user)
	})

	app.Post("/user", func(c *fiber.Ctx) error {
		var user User

		// Ambil body Json dan parsing ke struct
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Data tidak Valid",
			})
		}

		// Kembalikan data yang dikirim
		return c.JSON(fiber.Map{
			"message": "User diterima",
			"user":    user,
		})
	})

	app.Listen(":3000")
}
