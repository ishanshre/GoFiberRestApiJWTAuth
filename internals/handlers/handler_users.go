package handlers

import "github.com/gofiber/fiber/v2"

func (h *handler) AllUsers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
	})
}
