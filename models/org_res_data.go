package models

type ResOrg struct {
	Organisations []Organisation `json:"organisations"`
}

type OrgMessageData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type MultiOrgResData struct {
	OrgMessageData
	Data ResOrg `json:"data"`
}

type SingleOrgResData struct {
	OrgMessageData
	Data Organisation `json:"data"`
}

func NewOrgsResData(orgs []Organisation) MultiOrgResData {
	return MultiOrgResData{
		OrgMessageData{
			Status:  "success",
			Message: "organisations found",
		},
		ResOrg{Organisations: orgs},
	}
}

func NewOrgResData(org Organisation) SingleOrgResData {
	return SingleOrgResData{
		OrgMessageData{
			Status:  "success",
			Message: "organisations found",
		},
		org,
	}
}
