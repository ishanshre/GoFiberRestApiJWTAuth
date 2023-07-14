package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomInt(t *testing.T) {
	Init()

	min := int64(0)
	max := int64(100)

	result := RandomInt(min, max)
	assert.GreaterOrEqual(t, result, min)
	assert.LessOrEqual(t, result, max)
}

func TestRandomString(t *testing.T) {
	Init()

	length := 10
	result := RandomString(length)
	assert.Equal(t, length, len(result))
}

func TestRandomRole(t *testing.T) {
	Init()
	result := RandomRole()
	assert.GreaterOrEqual(t, result, 0)
	assert.LessOrEqual(t, result, 3)
}
