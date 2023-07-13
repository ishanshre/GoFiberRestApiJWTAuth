package dbrepo

import (
	"database/sql"

	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/config"
	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/repository"
)

type postgresDbRepo struct {
	Global *config.AppConfig
	DB     *sql.DB
}

func NewPostgresRepo(conn *sql.DB, global *config.AppConfig) repository.DatabaseRepo {
	return &postgresDbRepo{
		Global: global,
		DB:     conn,
	}
}
