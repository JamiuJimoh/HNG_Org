package models

type OrgReqData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AddUserToOrgReqData struct {
	UserId string `json:"userId"`
}
