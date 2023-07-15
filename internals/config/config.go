package config

import "log"

type AppConfig struct {
	InProduction bool
	Port         int
	DbString     string
	Dsn          string
	RedisHost    string
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
}
