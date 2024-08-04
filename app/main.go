package main

import (
	"ipl-api/database"
	"ipl-api/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	defer database.DB.Close()

	app.Use(logger.New())

	app.Get("/", handlers.HomepageHandler)
	app.Get("/team", handlers.TeamsHandler)
	app.Get("/player", handlers.PlayerHandler)
	app.Get("/league/batting", handlers.BattingRecordHandler)
	app.Get("/league/bowling", handlers.BowlingRecordHandler)

	log.Fatal(app.Listen(":8080"))
}
