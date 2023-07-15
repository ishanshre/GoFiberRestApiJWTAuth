package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/config"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/routers"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	var global config.AppConfig
	if err := godotenv.Load("../../.env"); err != nil {
		log.Printf("Error in loading environment files: %s\n", err.Error())
	}
	global.DbString = "postgres"
	global.Dsn = "test"
	testHandler, testDb, middleware := run(&global)
	defer testDb.SQL.Close()

	app := fiber.New()

	routers.Router(&global, app, testHandler, middleware)
	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/api/v1/users", nil))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
