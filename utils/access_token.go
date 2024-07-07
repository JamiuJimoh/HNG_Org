package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type AccessTokenConfig struct {
	secretKey []byte
}

func NewAccessToken(secretKey []byte) (*AccessTokenConfig, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}

	return &AccessTokenConfig{secretKey: secretKey}, nil
}

func (at *AccessTokenConfig) CreateToken(userID string, username string, duration time.Duration) (string, error) {
	claims, err := newTokenClaims(userID, username, duration)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(at.secretKey)
}

func (at *AccessTokenConfig) VerifyToken(tokenString []byte) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(string(tokenString), &TokenClaims{}, func(token *jwt.Token) (any, error) {
		return at.secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
