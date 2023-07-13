package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/config"
)

func Router(global *config.AppConfig, app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
}
