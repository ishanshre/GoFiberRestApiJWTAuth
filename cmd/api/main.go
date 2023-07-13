package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/config"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/drivers"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/handlers"
	dbrepo "github.com/ishanshre/GoFiberRestApiJWTAuth/internals/repository/dbRepo"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/routers"
	"github.com/joho/godotenv"
)

var global config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	handler := run()

	// create a new fiber app
	app := fiber.New()

	// pass fiber app to router to create routes
	routers.Router(&global, app, handler)

	// start the server
	app.Listen(fmt.Sprintf(":%d", global.Port))
}

func run() handlers.Handlers {
	// Configire flag and parse it.
	flag.IntVar(&global.Port, "port", 8000, "Port that servert listen to")
	flag.StringVar(&global.DbString, "dbString", "postgres", "Database string name")
	flag.Parse()

	// global config
	global.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	global.InfoLog = infoLog
	global.ErrorLog = errorLog

	// load environement files
	if err := godotenv.Load(".env"); err != nil {
		global.ErrorLog.Printf("Error in loading environment files: %s", err.Error())
	}

	// connect to database
	db, err := drivers.ConnectSql(global.DbString, os.Getenv(global.DbString))
	if err != nil {
		global.ErrorLog.Printf("error in connecting to database: %s", err.Error())
	}

	// closing the database connection at last
	defer db.SQL.Close()

	// connect to repository interface
	dbInterface := dbrepo.NewPostgresRepo(db.SQL, &global)

	handlerInterface := handlers.NewHandler(dbInterface, &global)

	return handlerInterface
}
