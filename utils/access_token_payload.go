package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	jwt.RegisteredClaims
}

func newTokenClaims(id string, username string, duration time.Duration) (*TokenClaims, error) {
	payload := &TokenClaims{
		ID:        id,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}
