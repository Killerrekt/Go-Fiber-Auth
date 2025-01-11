package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/killerrekt/Go-Fiber-Auth/utils"
)

func main() {
	app := fiber.New()

	utils.LoadConfig()

	fmt.Println(utils.Config)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})

	log.Fatal(app.Listen(":8080"))
}
