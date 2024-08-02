package handlers

import (
	"ipl-api/teams"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func TeamsHandler(ctx *fiber.Ctx) error {
	name, season := ctx.Query("name"), ctx.Query("season")

	if len(name) == 0 {
		return ctx.JSON(fiber.Map{
			"success": false,
			"body":    "please enter team name...",
		})
	}

	intSeason, err := strconv.Atoi(season)

	if err != nil {
		team, err := teams.GetAllTimeStats(name)
		if err != nil {
			return ctx.JSON(fiber.Map{
				"success": false,
				"body":    "failed to retrieve all-time stats...",
			})
		}
		return ctx.JSON(fiber.Map{
			"success": true,
			"body":    team,
		})
	}

	team, err := teams.GetSeasonStats(name, intSeason)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"success": false,
			"body":    "failed to retrive season stats...",
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"body":    team,
	})

}
