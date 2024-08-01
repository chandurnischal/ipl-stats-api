package handlers

import "github.com/gofiber/fiber/v2"

func HomepageHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"success": true,
		"body":    "Welcome to the IPL Stats API",
	})
}
