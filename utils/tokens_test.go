package utils

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func generateAccessToken(t *testing.T) *TokenDetail {
	id := int(RandomInt(10000000000, 10000000000000000))
	username := RandomString(1000)
	tokenID := uuid.NewV4().String()
	AccessTokenDetail, err := GenerateAccessToken(id, username, tokenID)
	assert.NoError(t, err)
	assert.NotNil(t, AccessTokenDetail)
	assert.NotEmpty(t, AccessTokenDetail)
	assert.NotZero(t, AccessTokenDetail.Token)
	assert.NotZero(t, AccessTokenDetail.TokenID)
	assert.NotZero(t, AccessTokenDetail.UserID)
	assert.NotZero(t, AccessTokenDetail.Username)
	assert.NotZero(t, AccessTokenDetail.ExpiresAt)
	assert.NotZero(t, AccessTokenDetail.Subject)
	assert.Equal(t, AccessTokenDetail.UserID, id)
	assert.Equal(t, AccessTokenDetail.Username, username)
	assert.Equal(t, AccessTokenDetail.Subject, "access_token")
	assert.Equal(t, AccessTokenDetail.TokenID, tokenID)
	return AccessTokenDetail
}

func TestGenerateAccessToken(t *testing.T) {
	generateAccessToken(t)
}

func TestGenerateRefreshToken(t *testing.T) {
	id := int(RandomInt(10000000000, 10000000000000000))
	username := RandomString(1000)
	tokenID := uuid.NewV4().String()
	RefreshTokenDetail, err := GenerateRefreshToken(id, username, tokenID)
	assert.NoError(t, err)
	assert.NotNil(t, RefreshTokenDetail)
	assert.NotEmpty(t, RefreshTokenDetail)
	assert.NotZero(t, RefreshTokenDetail.Token)
	assert.NotZero(t, RefreshTokenDetail.TokenID)
	assert.NotZero(t, RefreshTokenDetail.UserID)
	assert.NotZero(t, RefreshTokenDetail.Username)
	assert.NotZero(t, RefreshTokenDetail.ExpiresAt)
	assert.NotZero(t, RefreshTokenDetail.Subject)
	assert.Equal(t, RefreshTokenDetail.UserID, id)
	assert.Equal(t, RefreshTokenDetail.Username, username)
	assert.Equal(t, RefreshTokenDetail.Subject, "refresh_token")
	assert.Equal(t, RefreshTokenDetail.TokenID, tokenID)
}

func TestGenerateLoginResponse(t *testing.T) {
	id := int(RandomInt(10000000000, 10000000000000000))
	username := RandomString(1000)
	loginResponse, token, err := GenerateLoginResponse(id, username)
	assert.NoError(t, err)
	assert.NotNil(t, loginResponse)
	assert.NotNil(t, token)
	assert.NotEmpty(t, loginResponse)
	assert.NotEmpty(t, token)
	assert.NotZero(t, loginResponse.Username)
	assert.NotZero(t, loginResponse.ID)
	assert.NotZero(t, loginResponse.AccessToken)
	assert.NotZero(t, loginResponse.RefreshToken)
	assert.NotZero(t, token.RefreshToken.ExpiresAt)
	assert.NotZero(t, token.AccessToken.ExpiresAt)
	assert.Equal(t, loginResponse.ID, id)
	assert.Equal(t, loginResponse.Username, username)
	assert.Equal(t, token.AccessToken.TokenID, token.RefreshToken.TokenID)
	assert.Equal(t, token.AccessToken.UserID, id)
	assert.Equal(t, token.AccessToken.Username, username)
	assert.Equal(t, token.RefreshToken.Username, username)
	assert.Equal(t, token.RefreshToken.UserID, id)
	assert.Equal(t, token.RefreshToken.Subject, "refresh_token")
	assert.Equal(t, token.AccessToken.Subject, "access_token")
}
