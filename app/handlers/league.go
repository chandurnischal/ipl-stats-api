package handlers

import (
	"ipl-api/league"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func LeagueHandler(ctx *fiber.Ctx) error {
	season := ctx.Query("season")

	intSeason, err := strconv.Atoi(season)

	if err != nil {
		record := league.GetLeagueRecords(0)
		return ctx.JSON(fiber.Map{
			"success": true,
			"body":    record,
		})
	}

	record := league.GetLeagueRecords(intSeason)

	return ctx.JSON(fiber.Map{
		"success": true,
		"body":    record,
	})

}
