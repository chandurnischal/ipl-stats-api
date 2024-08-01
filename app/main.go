package main

import (
	"fmt"
	"ipl-api/batting"
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

	name, season := "virat kohli", 2024

	bat, err := batting.GetAllTimeBattingStats(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("All-time Stats:\n%v\n", bat)

	bat, err = batting.GetSeasonBattingStats(name, season)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nSeason %d Stats\n%v\n", season, bat)

	app.Use(logger.New())

	app.Get("/", handlers.HomepageHandler)
	app.Get("/team", handlers.TeamsHandler)

	// log.Fatal(app.Listen(":8080"))

}
