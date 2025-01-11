package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/killerrekt/Go-Fiber-Auth/db"
	"github.com/killerrekt/Go-Fiber-Auth/internal/dto/response"
	"github.com/killerrekt/Go-Fiber-Auth/internal/routes"
	"github.com/killerrekt/Go-Fiber-Auth/utils"
)

func main() {
	app := fiber.New()

	utils.LoadConfig()
	db.ConnectToDB()
	db.RunMigration()
	utils.SetUpValidator()
	routes.AuthRoute(app)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(response.Standard{
			Status:  false,
			Message: "Route not found",
		})
	})

	log.Fatal(app.Listen(":8080"))
}
