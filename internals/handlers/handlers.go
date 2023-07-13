package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/config"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/repository"
)

type Handlers interface {
	AllUsers(ctx *fiber.Ctx) error
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
