package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/config"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/repository"
)

type Handlers interface {
	AllUsers(ctx *fiber.Ctx) error
	GetUserByUsername(ctx *fiber.Ctx) error
	DeleteUserByUsername(ctx *fiber.Ctx) error
	RegisterUser(ctx *fiber.Ctx) error
	UsernameOrEmailExists(username, email string) (bool, error)
}

type handler struct {
	repo   repository.DatabaseRepo
	global *config.AppConfig
}

func NewHandler(r repository.DatabaseRepo, global *config.AppConfig) Handlers {
	return &handler{
		repo:   r,
		global: global,
	}
}
