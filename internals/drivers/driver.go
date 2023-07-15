package drivers

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	SQL *sql.DB
}

const (
	maxOpenDBConn = 10
	maxIdelDBConn = 5
	maxLifeDBTime = 5 * time.Minute
)

func ConnectSql(dbString, dsn string) (*DB, error) {
	d, err := newDatabase(dbString, dsn)
	if err != nil {
		return nil, err
	}

	d.SetMaxIdleConns(maxIdelDBConn)
	d.SetMaxOpenConns(maxOpenDBConn)
	d.SetConnMaxLifetime(maxLifeDBTime)

	return &DB{
		SQL: d,
	}, nil
}

func newDatabase(dbString, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbString, dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
