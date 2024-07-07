package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal json response with %v", payload)
		http.Error(w, "An error occured", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	type errResponse struct {
		Status     string `json:"status"`
		Message    string `json:"message"`
		StatusCode int    `json:"statusCode"`
	}
	RespondWithJSON(w, code, errResponse{
		Status:     "Bad request",
		Message:    message,
		StatusCode: code,
	})
}

// func RespondWithLoginData(w http.ResponseWriter, code int, message string, loginData models.LoginData) {
// 	RespondWithJSON(w, code, loginData{
// 		Status:  "success",
// 		Message: message,
// 		Data: resData{
// 			AccessToken: "",
// 			User:        resUser{},
// 		},
// 	})
// }
