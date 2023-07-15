package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/helpers"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/models"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/utils"
)

func (h *handler) RegisterUser(ctx *fiber.Ctx) error {
	userData := &models.ValidateUser{}
	if err := ctx.BodyParser(&userData); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(helpers.Message{
			Message: "error parsing data",
			Data:    err.Error(),
		})
	}

	if err := validate.Struct(userData); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(helpers.Message{
			MessageStatus: "error",
			Message:       err.Error(),
		})
	}

	exists, err := h.UsernameOrEmailExists(userData.Username, userData.Email)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(helpers.Message{
			MessageStatus: "error",
			Message:       err.Error(),
		})
	}
	if exists {
		return ctx.Status(http.StatusBadRequest).JSON(helpers.Message{
			MessageStatus: "error",
			Message:       "username/email already",
		})
	}
	hashedPassword, _ := utils.GeneratePassword(userData.Password)
	userData.Password = hashedPassword
	userData.CreatedAt = time.Now()
	userData.UpdatedAt = time.Now()
	user, err := h.repo.CreateUser(userData)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(helpers.Message{
			MessageStatus: "error",
			Message:       err.Error(),
		})
	}

	return ctx.Status(200).JSON(helpers.Message{
		MessageStatus: "success",
		Data:          user,
	})
}

func (h *handler) UserLogin(ctx *fiber.Ctx) error {
	user := &models.LoginUser{}
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(helpers.Message{
			MessageStatus: "error",
			Message:       fmt.Sprintf("error in parsing body: %s", err.Error()),
		})
	}
	exists, err := h.repo.UsernameExists(user.Username)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(helpers.Message{
			MessageStatus: "error",
			Message:       err.Error(),
		})
	}
	if !exists {
		return ctx.Status(http.StatusBadRequest).JSON(helpers.Message{
			MessageStatus: "error",
			Message:       "username does not exists",
		})
	}
	userData, err := h.repo.GetUserByUsername(user.Username)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(helpers.Message{
			MessageStatus: "error",
			Message:       err.Error(),
		})
	}
	if err := utils.ComparePassword(userData.Password, user.Password); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(helpers.Message{
			MessageStatus: "error",
			Message:       "password does not match",
		})
	}
	loginResponse, tokens, err := utils.GenerateLoginResponse(userData.ID, userData.Username)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(helpers.Message{
			MessageStatus: "error",
			Message:       err.Error(),
		})
	}
	tokensJSON, _ := json.Marshal(tokens)
	if err := h.redisHost.Set(context.Background(), tokens.AccessToken.TokenID, tokensJSON, time.Until(tokens.AccessToken.ExpiresAt)).Err(); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(helpers.Message{
			MessageStatus: "error",
			Message:       err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(helpers.Message{
		MessageStatus: "success",
		Data:          loginResponse,
	})
}
