package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/killerrekt/Go-Fiber-Auth/db"
	"github.com/killerrekt/Go-Fiber-Auth/utils"
)

func main() {
	app := fiber.New()

	utils.LoadConfig()
	db.ConnectToDB()
	db.RunMigration()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})

	log.Fatal(app.Listen(":8080"))
}
