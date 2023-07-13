package config

import "log"

type AppConfig struct {
	InProduction bool
	Port         int
	DbDriverName string
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
}
