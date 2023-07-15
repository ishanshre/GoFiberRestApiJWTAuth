package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/config"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/drivers"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/handlers"
	dbrepo "github.com/ishanshre/GoFiberRestApiJWTAuth/internals/repository/dbRepo"
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

	h = handlers.NewHandler(dbInterface, global)
	Router(global, app, h)
	res, err := app.Test(httptest.NewRequest(http.MethodGet, "/api/v1/users", nil))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
