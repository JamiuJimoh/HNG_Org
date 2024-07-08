package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/JamiuJimoh/hngorg/utils"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	currentUserIDKey        = "currentuseridkey"
)

func (ac *ApiCfg) AuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		bearerToken, err := getBearerToken(r.Header)
		claims, err := ac.tokenCfg.VerifyToken([]byte(bearerToken))
		if err != nil {
			handleLoginError(w, err)
			return
		}
		ac.ctx = context.WithValue(ac.ctx, currentUserIDKey, claims.ID)

		// user, err := ac.db.GetUserByID(r.Context(), claims.ID)
		// if err != nil {
		// 	handleLoginError(w, err)
		// 	return
		// }
		// log.Print(user)
		handler(w, r)
	}
}

func getBearerToken(h http.Header) (string, error) {
	authorizationHeader := h.Get(authorizationHeaderKey)
	if len(authorizationHeader) == 0 {
		err := errors.New("authentication failed: authorization header is not provided")

		return "", err
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) != 2 {
		err := errors.New("authentication failed: invalid authorization header format")
		return "", err
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationTypeBearer {
		err := errors.New("authentication failed: malformed authorizationa type")
		return "", err
	}

	return fields[1], nil
}

func handleLoginError(w http.ResponseWriter, err error) {
	log.Print(err)
	utils.RespondWithError(w, http.StatusUnauthorized, "Authentication failed")
}

func handleAuthorizationError(w http.ResponseWriter, err error) {
	log.Print(err)
	utils.RespondWithError(w, http.StatusBadRequest, "Client error")
}
