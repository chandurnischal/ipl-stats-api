package handlers

import (
	"ipl-api/players"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PlayerHandler(ctx *fiber.Ctx) error {
	name := ctx.Query("name")
	season := ctx.Query("season")

	if season == "" {
		playerCard, err := players.GetAllTimePlayerCard(name)
		if err != nil {
			return ctx.JSON(fiber.Map{
				"success": false,
				"body":    "failed to retrieve all-time stats...",
			})
		}
		return ctx.JSON(fiber.Map{
			"success": true,
			"body":    playerCard,
		})
	}

	intSeason, err := strconv.Atoi(season)

	if err != nil {
		return ctx.JSON(fiber.Map{
			"success": false,
			"body":    "incorrect season...",
		})
	}

	playerCard, err := players.GetSeasonPlayerCard(name, intSeason)

	if err != nil {
		ctx.JSON(fiber.Map{
			"success": false,
			"body":    "failed to retrieve season stats...",
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"body":    playerCard,
	})

}
