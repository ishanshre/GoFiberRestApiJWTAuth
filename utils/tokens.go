package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
)

var (
	AccessExpiresAt  = jwt.NewNumericDate(time.Now().Add(time.Minute * 15))
	RefreshExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
	IssuedAt         = jwt.NewNumericDate(time.Now())
	NotBefore        = jwt.NewNumericDate(time.Now())
	Secret           = []byte(os.Getenv("jwt_secret"))
)

type Claims struct {
	Username string
	ID       int
	jwt.RegisteredClaims
}

type TokenDetail struct {
	Token     *string
	TokenID   string
	UserID    int
	Username  string
	ExpiresAt time.Time
	Subject   string
}

type Token struct {
	AccessToken  *TokenDetail
	RefreshToken *TokenDetail
}

type LoginResponse struct {
	Username     string `json:"username"`
	ID           int    `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateLoginResponse(id int, username string) (*LoginResponse, *Token, error) {
	tokenID := uuid.NewV4().String()
	access_token_detail, err := GenerateAccessToken(id, username, tokenID)
	if err != nil {
		return nil, nil, err
	}
	refresh_token_detail, err := GenerateRefreshToken(id, username, tokenID)
	if err != nil {
		return nil, nil, err
	}
	return &LoginResponse{
			Username:     username,
			ID:           id,
			AccessToken:  *access_token_detail.Token,
			RefreshToken: *refresh_token_detail.Token,
		}, &Token{
			AccessToken:  access_token_detail,
			RefreshToken: refresh_token_detail,
		}, nil
}

func GenerateAccessToken(id int, username, tokenID string) (*TokenDetail, error) {
	// create claims for access token
	access_claims := &Claims{
		Username: username,
		ID:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: AccessExpiresAt,
			IssuedAt:  IssuedAt,
			NotBefore: NotBefore,
			Subject:   "access_token",
			ID:        tokenID,
		},
	}

	// creating a new token access claims and signing method
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, access_claims)

	// sign the token with the unique secret key from the env files
	signedAccessToken, err := access_token.SignedString(Secret)
	if err != nil {
		return nil, err
	}
	return &TokenDetail{
		Token:     &signedAccessToken,
		UserID:    access_claims.ID,
		Username:  access_claims.Username,
		ExpiresAt: access_claims.RegisteredClaims.ExpiresAt.Time,
		Subject:   access_claims.Subject,
		TokenID:   access_claims.RegisteredClaims.ID,
	}, nil
}

func GenerateRefreshToken(id int, username, tokenID string) (*TokenDetail, error) {
	// create claims for refresh token
	refresh_claims := &Claims{
		ID:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: RefreshExpiresAt,
			Subject:   "refresh_token",
			IssuedAt:  IssuedAt,
			NotBefore: NotBefore,
			ID:        tokenID,
		},
	}

	// create token with the claims
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims)

	// sign the token with the secret
	signedRefreshToken, err := refresh_token.SignedString(Secret)
	if err != nil {
		return nil, err
	}

	return &TokenDetail{
		Token:     &signedRefreshToken,
		TokenID:   refresh_claims.RegisteredClaims.ID,
		UserID:    refresh_claims.ID,
		Username:  refresh_claims.Username,
		ExpiresAt: refresh_claims.RegisteredClaims.ExpiresAt.Time,
		Subject:   refresh_claims.RegisteredClaims.Subject,
	}, err

}
