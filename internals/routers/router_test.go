package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/config"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/drivers"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/handlers"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/middlewares"
	dbrepo "github.com/ishanshre/GoFiberRestApiJWTAuth/internals/repository/dbRepo"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	global := &config.AppConfig{}
	app := fiber.New()
	var h handlers.Handlers
	db, err := drivers.ConnectSql(dbString, dsn)
	if err != nil {
		t.Errorf("Cannot connect to database: %v", err)
	}

	dbInterface := dbrepo.NewPostgresRepo(db.SQL, global)
	redisPool := redis.NewClient(
		&redis.Options{
			Addr:         global.RedisHost,
			Password:     "",
			DB:           0,
			MaxIdleConns: 10,
		},
	)

	h = handlers.NewHandler(dbInterface, global, redisPool)
	m := middlewares.NewMiddleware(redisPool)
	Router(global, app, h, m)
	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/api/v1/users", nil))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
