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

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Post("/login", h.UserLogin)

	user := v1.Group("/users")
	user.Get("/", h.AllUsers)
	user.Post("/", h.RegisterUser)
	user.Get("/:username", h.GetUserByUsername)
	user.Delete("/:username/delete", h.DeleteUserByUsername)
}
