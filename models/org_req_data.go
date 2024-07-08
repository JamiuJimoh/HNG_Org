package models

type OrgReqData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type OrgMemberReqData struct {
	UserId string `json:"userId"`
}
