package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/config"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/handlers"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/middlewares"
)

func Router(global *config.AppConfig, app *fiber.App, h handlers.Handlers, m middlewares.MiddlewareRepo) {
	app.Use(cors.New())

	app.Use(logger.New())
	app.Use(helmet.New())
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Post("/login", h.UserLogin)
	v1.Post("/", h.RegisterUser)
	user := v1.Group("/users", m.JwtAuth())
	user.Get("/", h.AllUsers)
	user.Get("/:username", h.GetUserByUsername)
	user.Delete("/:username/delete", h.DeleteUserByUsername)
	user.Post("/:username/logout", h.UserLogout)
}
