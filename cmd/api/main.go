package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/config"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/routers"
)

var global config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	flag.IntVar(&global.Port, "port", 8000, "Port that servert listen to")
	flag.StringVar(&global.DbString, "dbString", "postgres", "Database string name")
	flag.Parse()
	app := fiber.New()
	routers.Router(&global, app)
	app.Listen(fmt.Sprintf(":%d", global.Port))
}
