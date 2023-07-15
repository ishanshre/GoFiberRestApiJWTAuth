package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type MiddlewareRepo interface {
	JwtAuth() fiber.Handler
}

type middlewares struct {
	redisClient *redis.Client
}

func NewMiddleware(r *redis.Client) MiddlewareRepo {
	return &middlewares{
		redisClient: r,
	}
}
