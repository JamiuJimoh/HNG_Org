package models

import (
	"github.com/JamiuJimoh/hngorg/db/sqlc"
)

type Organisation struct {
	OrgId       string `json:"orgId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func OrgFromSQLOrg(org sqlc.Organisation) Organisation {
	return Organisation{
		OrgId:       org.OrgID,
		Name:        org.Name,
		Description: org.Description.String,
	}
}

func OrgsFromSQLOrgs(orgs []sqlc.Organisation) []Organisation {
	var organisations []Organisation
	for _, org := range orgs {
		organisations = append(organisations, OrgFromSQLOrg(org))
	}
	return organisations
}
