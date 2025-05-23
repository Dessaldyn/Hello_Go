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

	app.Listen(":3000")
}
