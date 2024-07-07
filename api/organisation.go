package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/JamiuJimoh/hngorg/db/sqlc"
	"github.com/JamiuJimoh/hngorg/models"
	"github.com/JamiuJimoh/hngorg/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (ac *ApiCfg) GetOrganistions(w http.ResponseWriter, r *http.Request) {
	currentUserId := (ac.ctx.Value(currentUserIDKey)).(string)
	orgs, err := ac.db.GetUserOrgsByID(r.Context(), currentUserId)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "no results found")
		return
	}

	data := models.OrgsFromSQLOrgs(orgs)
	utils.RespondWithJSON(w, http.StatusOK, models.NewOrgsResData(data))
}

func (ac *ApiCfg) GetOrganistion(w http.ResponseWriter, r *http.Request) {
	orgId := r.PathValue("orgId")
	org, err := ac.db.GetOrgByID(r.Context(), orgId)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "no results found")
		return
	}

	data := models.OrgFromSQLOrg(org)
	utils.RespondWithJSON(w, http.StatusOK, models.NewOrgResData(data))
}

func (ac *ApiCfg) CreateOrganistion(w http.ResponseWriter, r *http.Request) {
	currentUserId := (ac.ctx.Value(currentUserIDKey)).(string)
	var org models.OrgReqData
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "an error occurred")
		return
	}

	err = json.Unmarshal(bytes, &org)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Client error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	validatorError := utils.ValidateOrgName(org.Name)
	if validatorError != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(validatorError)
		return
	}

	id, err := uuid.NewV7()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "an error occurred")
		return
	}
	sqlOrg, err := ac.db.CreateOrganisation(r.Context(), sqlc.CreateOrganisationParams{
		OrgID:       id.String(),
		Name:        org.Name,
		Description: pgtype.Text{String: org.Description, Valid: true},
		UserID:      currentUserId,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Client error")
		return
	}

	ret := models.OrgFromSQLOrg(sqlOrg)
	utils.RespondWithJSON(w, http.StatusCreated, ret)
}

func (ac *ApiCfg) PatchOrganistionWithUser(w http.ResponseWriter, r *http.Request) {
	// currentUserId := (ac.ctx.Value(currentUserIDKey)).(string)
	// orgId := r.PathValue("orgId")
	// sqlOrg, err := ac.db.GetOrgByOrgID(r.Context(), orgId)
	// if err != nil {
	// utils.RespondWithError(w, http.StatusNotFound, "Client error")
	// return
	// }
}
