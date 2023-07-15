package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpperCase(t *testing.T) {
	type TestStruct struct {
		Field1 string `validate:"upper"`
	}

	validStruct := TestStruct{
		Field1: "Abc112",
	}
	err := validate.Struct(validStruct)
	assert.NoError(t, err)

	// invalid test
	inValidStruct := TestStruct{
		Field1: "asdasdasd",
	}
	err = validate.Struct(inValidStruct)
	assert.Error(t, err)

}

func TestLowerCase(t *testing.T) {
	type TestStruct struct {
		Field1 string `validate:"lower"`
	}

	validStruct := TestStruct{
		Field1: "abc112",
	}
	err := validate.Struct(validStruct)
	assert.NoError(t, err)

	// invalid test
	inValidStruct := TestStruct{
		Field1: "AAAAAA",
	}
	err = validate.Struct(inValidStruct)
	assert.Error(t, err)

}

func TestNumber(t *testing.T) {
	type TestStruct struct {
		Field1 string `validate:"number"`
	}

	validStruct := TestStruct{
		Field1: "123234sdad",
	}
	err := validate.Struct(validStruct)
	assert.NoError(t, err)

	// invalid test
	inValidStruct := TestStruct{
		Field1: "asdasdAsdf",
	}
	err = validate.Struct(inValidStruct)
	assert.Error(t, err)

}
