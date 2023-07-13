package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/config"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/drivers"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/routers"
	"github.com/joho/godotenv"
)

var global config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
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

	db, err := run()
	if err != nil {
		global.ErrorLog.Println(err)
	}

	defer db.SQL.Close()

	// create a new fiber app
	app := fiber.New()

	// pass fiber app to router to create routes
	routers.Router(&global, app)

	// start the server
	app.Listen(fmt.Sprintf(":%d", global.Port))
}

func run() (*drivers.DB, error) {

	// load environement files
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("error in loading the environment files: %v", err.Error())
	}

	// connect to database
	db, err := drivers.ConnectSql(global.DbString, os.Getenv(global.DbString))
	if err != nil {
		return nil, fmt.Errorf("error in connecting to database: %v", err.Error())
	}
	return db, nil
}
