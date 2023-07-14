package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func generatePassword(t *testing.T) (string, string) {
	password := RandomString(20)
	hashedPassword, err := GeneratePassword(password)
	assert.NoError(t, err)
	return hashedPassword, password
}

func TestGeneratePassword(t *testing.T) {
	hashedPassword, _ := generatePassword(t)
	if hashedPassword == "" {
		t.Errorf("expected a hashedPassword, but got %v", hashedPassword)
	}
}

func TestGeneratePasswordFailure(t *testing.T) {
	password := RandomString(100)
	hashedPassword, err := GeneratePassword(password)
	assert.Error(t, err)
	assert.EqualError(t, err, bcrypt.ErrPasswordTooLong.Error())
	assert.Equal(t, hashedPassword, "")
}

func TestComparePassword(t *testing.T) {
	hashedPassword, plainPassword := generatePassword(t)
	err := ComparePassword(hashedPassword, plainPassword)
	assert.NoError(t, err)
}
