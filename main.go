package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sixfwa/fiber-api/database"
	"github.com/sixfwa/fiber-api/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}

func setupRoutes(app *fiber.App) {
	// welcome
	app.Get("/api", welcome)
	// Year endpoints
	app.Post("/api/years", routes.CreateYear)
	app.Get("/api/years", routes.GetYears)
	app.Get("/api/years/:id", routes.GetYear)
	app.Put("/api/years/:id", routes.UpdateYear)
	app.Delete("/api/years/:id", routes.DeleteYear)
	// Item endpoints
	app.Post("/api/items", routes.CreateItem)
	app.Get("/api/items", routes.GetItems)
	app.Get("/api/items/:id", routes.GetItem)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)
	log.Fatal(app.Listen(":3005"))
}
