package api

import (
	"context"
	"os"

	"github.com/JamiuJimoh/hngorg/db/sqlc"
	"github.com/JamiuJimoh/hngorg/utils"
)

type ApiCfg struct {
	db       *sqlc.Queries
	tokenCfg *utils.AccessTokenConfig
	ctx      context.Context
}

func NewApiConfig(db *sqlc.Queries, tokenCfg *utils.AccessTokenConfig) (*ApiCfg, error) {
	secretKey := os.Getenv("TOKEN_SYMMETRIC_KEY")
	tokenCfg, err := utils.NewAccessToken([]byte(secretKey))
	if err != nil {
		return nil, err
	}
	cfg := &ApiCfg{db: db,
		tokenCfg: tokenCfg,
		ctx:      context.Background(),
	}
	return cfg, nil
}
