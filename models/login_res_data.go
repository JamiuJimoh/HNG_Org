package models

type LoginRes struct {
	AccessToken string  `json:"accessToken"`
	User        ResUser `json:"user"`
}
type LoginResData struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    LoginRes `json:"data"`
}
