package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/helpers"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/utils"
)

func (m *middlewares) JwtAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// get the auth token from the header
		bearerToken := ctx.Get("Authorization")

		// check if the bearer token is empty. If empty return error
		if bearerToken == "" {
			return ctx.Status(http.StatusUnauthorized).JSON(helpers.Message{
				MessageStatus: "error",
				Message:       "authentication token not provided",
			})
		}

		// check the auth token  format
		tokenString := strings.Split(bearerToken, " ")
		if len(tokenString) != 2 && tokenString[0] != "Bearer" {
			return ctx.Status(http.StatusUnauthorized).JSON(helpers.Message{
				MessageStatus: "error",
				Message:       "invalid authentication token format",
			})
		}

		// verify the token
		tokenDetail, err := utils.VerifyTokenWithClaims(tokenString[1], "access_token")
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(helpers.Message{
				MessageStatus: "error",
				Message:       fmt.Sprintf("invalid token: %s", err.Error()),
			})
		}
		if err := m.redisClient.Exists(context.Background(), tokenDetail.TokenID).Err(); err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(helpers.Message{
				MessageStatus: "error",
				Message:       "Invalid token or token does not exists in cache",
			})
		}
		ctx.Locals("tokenID", tokenDetail.TokenID)
		ctx.Locals("userID", tokenDetail.UserID)
		ctx.Locals("username", tokenDetail.Username)
		return ctx.Next()
	}
}
