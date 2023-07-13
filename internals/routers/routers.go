package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/config"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/handlers"
)

func Router(global *config.AppConfig, app *fiber.App, h handlers.Handlers) {
	app.Use(cors.New())

	app.Use(logger.New())

	app.Get("/", h.AllUsers)
}
