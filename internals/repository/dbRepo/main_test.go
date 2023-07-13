package dbrepo

import (
	"log"
	"os"
	"testing"

	"github.com/ishanshre/GoFiberRestApiJWTAuth/internals/drivers"
	"github.com/joho/godotenv"
)

var testQueries *postgresDbRepo

func TestMain(m *testing.M) {
	if err := godotenv.Load("./../../../.env"); err != nil {
		log.Fatalf("cannot load env files: %v", err)
	}
	db, err := drivers.ConnectSql("postgres", os.Getenv("test"))
	if err != nil {
		log.Fatalf("error in connecting to test database: %v", err)
	}
	defer db.SQL.Close()
	testQueries = &postgresDbRepo{
		DB: db.SQL,
	}
	os.Exit(m.Run())
}
