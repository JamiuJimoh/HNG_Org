package models

type ResOrg struct {
	Organisations []Organisation `json:"organisations"`
}

type MultiOrgResData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    ResOrg `json:"data"`
}

type SingleOrgResData struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    Organisation `json:"data"`
}

func NewOrgsResData(orgs []Organisation) MultiOrgResData {
	return MultiOrgResData{
		Status:  "success",
		Message: "organisations found",
		Data: ResOrg{
			Organisations: orgs,
		},
	}
}

func NewOrgResData(org Organisation) SingleOrgResData {
	return SingleOrgResData{
		Status:  "success",
		Message: "organisation found",
		Data:    org,
	}
}
