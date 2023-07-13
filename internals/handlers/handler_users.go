package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/helpers"
)

func (h *handler) AllUsers(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 10)
	offset := ctx.QueryInt("offset", 0)
	log.Println(time.Now())
	users, err := h.repo.AllUsers(limit, offset)
	if err != nil {
		log.Println(err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return ctx.Status(200).JSON(helpers.Message{
		Message: "Success",
		Limit:   limit,
		Offset:  offset,
		Data:    users,
	})
}
