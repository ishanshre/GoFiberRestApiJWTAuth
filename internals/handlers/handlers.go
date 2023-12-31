package handlers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/config"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/repository"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/validators"
	"github.com/redis/go-redis/v9"
)

type Handlers interface {
	AllUsers(ctx *fiber.Ctx) error
	GetUserByUsername(ctx *fiber.Ctx) error
	DeleteUserByUsername(ctx *fiber.Ctx) error
	RegisterUser(ctx *fiber.Ctx) error
	UsernameOrEmailExists(username, email string) (bool, error)
	UserLogin(ctx *fiber.Ctx) error
	UserLogout(ctx *fiber.Ctx) error
	Refresh(ctx *fiber.Ctx) error
}

type handler struct {
	repo      repository.DatabaseRepo
	global    *config.AppConfig
	redisHost *redis.Client
}

var validate *validator.Validate

func NewHandler(r repository.DatabaseRepo, global *config.AppConfig, rH *redis.Client) Handlers {
	validate = validator.New()
	validate.RegisterValidation("upper", validators.UpperCase)
	validate.RegisterValidation("lower", validators.LowerCase)
	validate.RegisterValidation("number", validators.Number)
	return &handler{
		repo:      r,
		global:    global,
		redisHost: rH,
	}
}
