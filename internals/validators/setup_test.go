package validators

import (
	"os"
	"testing"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

func TestMain(m *testing.M) {
	validate = validator.New()
	validate.RegisterValidation("upper", UpperCase)
	validate.RegisterValidation("lower", LowerCase)
	validate.RegisterValidation("number", Number)
	os.Exit(m.Run())
}
