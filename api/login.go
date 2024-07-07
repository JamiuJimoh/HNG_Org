package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/JamiuJimoh/hngorg/models"
	"github.com/JamiuJimoh/hngorg/utils"
	"golang.org/x/crypto/bcrypt"
)

func (ac *ApiCfg) Login(w http.ResponseWriter, r *http.Request) {
	type LoginReqData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var loginData LoginReqData
	read, err := io.ReadAll(r.Body)
	if err != nil {
		handleLoginError(w, err)
		return
	}
	json.Unmarshal(read, &loginData)

	// use database stored hash by finding user using email
	dbUser, err := ac.db.GetUserByEmail(r.Context(), loginData.Email)
	if err != nil {
		handleLoginError(w, err)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginData.Password))
	if err != nil {
		handleLoginError(w, err)
		return
	}
	payload, err := ac.createLoginData(models.ResUserFromSQLUser(dbUser), "Login successful")
	if err != nil {
		handleLoginError(w, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, payload)
}

func (ac *ApiCfg) createLoginData(u models.ResUser, message string) (*models.LoginResData, error) {
	accessToken, err := ac.tokenCfg.CreateToken(u.UserId, u.FirstName, time.Duration(time.Minute*15))
	if err != nil {
		return nil, err
	}

	loginData := &models.LoginResData{
		Status:  "success",
		Message: message,
		Data: models.LoginRes{
			AccessToken: accessToken,
			User:        u,
		},
	}
	return loginData, nil
}
