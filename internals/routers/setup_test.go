package routers

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var dbString string
var dsn string

func TestMain(m *testing.M) {
	if err := godotenv.Load("./../../.env"); err != nil {
		log.Fatalf("cannot load env files: %v", err)
	}
	dbString = "postgres"
	dsn = os.Getenv("test")

	os.Exit(m.Run())
}
