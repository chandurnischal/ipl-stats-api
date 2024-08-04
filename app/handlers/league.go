package handlers

import (
	"ipl-api/league"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func BattingRecordHandler(ctx *fiber.Ctx) error {
	season := ctx.Query("season")

	intSeason, err := strconv.Atoi(season)

	if err != nil {
		record := league.GetBattingRecord(0)
		return ctx.JSON(fiber.Map{
			"success": true,
			"body":    record,
		})
	}

	record := league.GetBattingRecord(intSeason)

	return ctx.JSON(fiber.Map{
		"success": true,
		"body":    record,
	})
}

func BowlingRecordHandler(ctx *fiber.Ctx) error {
	season := ctx.Query("season")

	intSeason, err := strconv.Atoi(season)

	if err != nil {
		record := league.GetBowlingRecord(0)
		return ctx.JSON(fiber.Map{
			"success": true,
			"body":    record,
		})
	}

	record := league.GetBowlingRecord(intSeason)

	return ctx.JSON(fiber.Map{
		"success": true,
		"body":    record,
	})
}
