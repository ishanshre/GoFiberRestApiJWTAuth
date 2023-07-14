package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/helpers"
)

func (h *handler) AllUsers(ctx *fiber.Ctx) error {
	limit := ctx.QueryInt("limit", 10)
	offset := ctx.QueryInt("offset", 0)
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

func (h *handler) GetUserByUsername(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	user, err := h.repo.GetUserByUsername(username)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(helpers.Message{
			Message: "error",
			Data:    err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(helpers.Message{
		Message: "success",
		Data:    user,
	})
}

func (h *handler) DeleteUserByUsername(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	if err := h.repo.DeleteUser(username); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(helpers.Message{
			Message: "error",
			Data:    err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(helpers.Message{
		Message: "success",
	})

}
