package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/config"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/drivers"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/handlers"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/middlewares"
	dbrepo "github.com/ishanshre/GoFiberRestApiJWTAuth/internals/repository/dbRepo"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/routers"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	var global config.AppConfig
	// load environement files
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error in loading environment files: %s\n", err.Error())
	}

	// Configire flag and parse it.
	flag.IntVar(&global.Port, "port", 8000, "Port that servert listen to")
	flag.StringVar(&global.DbString, "dbString", "postgres", "Database string name")
	flag.StringVar(&global.Dsn, "dsn", "postgres", "maps to env key pair")
	flag.StringVar(&global.RedisHost, "redisHost", "redis:6379", "address to redis")
	flag.Parse()

	handler, db, middleware := run(&global)

	// closing the database connection at last
	defer db.SQL.Close()

	// create a new fiber app
	app := fiber.New()

	// pass fiber app to router to create routes
	routers.Router(&global, app, handler, middleware)

	// start the server
	app.Listen(fmt.Sprintf(":%d", global.Port))
}

func run(global *config.AppConfig) (handlers.Handlers, *drivers.DB, middlewares.MiddlewareRepo) {

	// global config
	global.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	global.InfoLog = infoLog
	global.ErrorLog = errorLog

	log.Println(os.Getenv("postgres"))
	log.Println(os.Getenv("DB_URL"))
	log.Println(os.Getenv("test"))

	// connect to database
	db, err := drivers.ConnectSql(global.DbString, os.Getenv("DB_URL"))
	if err != nil {
		global.ErrorLog.Printf("error in connecting to database: %s", err.Error())
	}

	redisPool := redis.NewClient(
		&redis.Options{
			Addr:         os.Getenv("REDIS_URL"),
			Password:     "",
			DB:           0,
			MaxIdleConns: 10,
			PoolSize:     10,
			MinIdleConns: 0,
		},
	)

	if err := redisPool.Ping(context.Background()).Err(); err != nil {
		global.ErrorLog.Printf("error in connecting to redis: %s", err.Error())
	}

	// connect to repository interface
	dbInterface := dbrepo.NewPostgresRepo(db.SQL, global)

	handlerInterface := handlers.NewHandler(dbInterface, global, redisPool)

	middleware := middlewares.NewMiddleware(redisPool)

	return handlerInterface, db, middleware
}
