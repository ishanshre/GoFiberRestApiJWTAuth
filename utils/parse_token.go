package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyTokenWithClaims(tokenString, subject string) (*TokenDetail, error) {
	claims := &Claims{}
	token, err := ExtractToken(tokenString, subject, claims)
	if err != nil {
		return nil, err
	}
	if err := ValidateToken(token, claims, subject); err != nil {
		return nil, err
	}
	return &TokenDetail{
		Token:     &tokenString,
		TokenID:   claims.RegisteredClaims.ID,
		UserID:    claims.ID,
		Username:  claims.Username,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Time,
		Subject:   claims.RegisteredClaims.Subject,
	}, nil
}

func ExtractToken(tokenString, subject string, claims *Claims) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalud token signing method")
		}
		return Secret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("token signature invalid: %v", err)
		}
		return nil, err
	}
	return token, nil
}

func ValidateToken(token *jwt.Token, claims *Claims, subject string) error {
	if !token.Valid {
		return errors.New("token is not valid")
	}
	if time.Now().Unix() > claims.RegisteredClaims.ExpiresAt.Unix() {
		return errors.New("token already expired")
	}
	if claims.RegisteredClaims.Subject != subject {
		return errors.New("token subject mismatch")
	}
	return nil
}
