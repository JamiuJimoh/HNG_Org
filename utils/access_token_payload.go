package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func newTokenClaims(id string, username string, duration time.Duration) (*TokenClaims, error) {
	payload := &TokenClaims{
		id,
		username,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
	return payload, nil
}
