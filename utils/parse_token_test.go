package utils

import (
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVerifyTokenWithClaims(t *testing.T) {
	tokenDetail := generateAccessToken(t)
	assert.NotNil(t, tokenDetail)
	assert.NotEmpty(t, tokenDetail)
	tokenDetailValidated, err := VerifyTokenWithClaims(*tokenDetail.Token, tokenDetail.Subject)
	assert.NoError(t, err)
	assert.NotNil(t, tokenDetailValidated)
}

func TestVerifyTokenWithClaims_Failure(t *testing.T) {
	tokenString, err := generateInvalidTokenWithClaims(t)
	assert.NoError(t, err)
	require.NotEmpty(t, tokenString)
	tokenDetail, err := VerifyTokenWithClaims(tokenString, "access_token")
	assert.ErrorContains(t, err, jwt.ErrTokenInvalidClaims.Error())
	assert.Nil(t, tokenDetail)

	id := 1
	username := RandomString(6)
	tokenID := uuid.NewV4().String()
	accessTokenDetail, err := GenerateAccessToken(id, username, tokenID)
	assert.NoError(t, err)
	assert.NotNil(t, accessTokenDetail)
	tokenDetail, err = VerifyTokenWithClaims(*accessTokenDetail.Token, "refresh_token")
	assert.Error(t, err)
	assert.ErrorContains(t, err, "token subject mismatch")
	assert.Empty(t, tokenDetail)
}

func TestExtractToken_Failure(t *testing.T) {
	invalidToken := "invalid-token"
	claims := &Claims{}
	token, err := ExtractToken(invalidToken, "token", claims)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, jwt.ErrTokenMalformed))
	assert.Empty(t, token)
}

func generateInvalidTokenWithClaims(t *testing.T) (string, error) {
	username := RandomString(10)
	id := int(RandomInt(1000000000, 1000000000000))
	tokenID := uuid.NewV4().String()
	jwtExpire := jwt.NewNumericDate(time.Now().Add(-5 * time.Minute))
	jwtNow := jwt.NewNumericDate(time.Now())
	claims := &Claims{
		Username: username,
		ID:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwtExpire,
			IssuedAt:  jwtNow,
			NotBefore: jwtNow,
			Subject:   "access_token",
			ID:        tokenID,
		},
	}

	// creating a new token access claims and signing method
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the token with the unique secret key from the env files
	signedAccessToken, err := access_token.SignedString(Secret)
	if err != nil {
		return "", err
	}
	return signedAccessToken, nil
}
