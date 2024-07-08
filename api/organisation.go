package api

import (
	"encoding/json"
	"io"
	"log"
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
		log.Print(err)
		utils.RespondWithError(w, http.StatusNotFound, "no results found")
		return
	}

	data := models.OrgsFromSQLOrgs(orgs)
	utils.RespondWithJSON(w, http.StatusOK, models.FoundOrgsResData(data))
}

func (ac *ApiCfg) GetOrganistion(w http.ResponseWriter, r *http.Request) {
	orgId := r.PathValue("orgId")
	org, err := ac.db.GetOrgByID(r.Context(), orgId)
	if err != nil {
		log.Print(err)
		utils.RespondWithError(w, http.StatusNotFound, "no results found")
		return
	}

	data := models.OrgFromSQLOrg(org)
	utils.RespondWithJSON(w, http.StatusOK, models.FoundOrgResData(data))
}

func (ac *ApiCfg) CreateOrganistion(w http.ResponseWriter, r *http.Request) {
	currentUserId := (ac.ctx.Value(currentUserIDKey)).(string)
	var org models.OrgReqData
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		utils.RespondWithError(w, http.StatusInternalServerError, "an error occurred")
		return
	}

	err = json.Unmarshal(bytes, &org)
	if err != nil {
		log.Print(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Client error")
		return
	}

	validatorError := utils.ValidateOrgName(org.Name)
	if validatorError != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(validatorError)
		return
	}

	id, err := uuid.NewV7()
	if err != nil {
		log.Print(err)
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
		log.Print(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Client error")
		return
	}

	_, err = ac.db.CreateOrgMember(r.Context(), sqlc.CreateOrgMemberParams{
		MemberID:  sqlOrg.UserID,
		OrgID:     sqlOrg.OrgID,
		CreatorID: sqlOrg.UserID,
	})
	if err != nil {
		handleRegistrationError(w, err)
		return
	}

	data := models.OrgFromSQLOrg(sqlOrg)
	utils.RespondWithJSON(w, http.StatusCreated, models.NewOrgResData(data))
}

func (ac *ApiCfg) PatchOrganistionWithUser(w http.ResponseWriter, r *http.Request) {
	orgId := r.PathValue("orgId")

	var member models.OrgMemberReqData
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		utils.RespondWithError(w, http.StatusInternalServerError, "an error occurred")
		return
	}
	err = json.Unmarshal(bytes, &member)
	if err != nil {
		log.Print(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Client error")
		return
	}

	sqlOrg, err := ac.db.GetOrgByID(r.Context(), orgId)
	if err != nil {
		log.Print(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Client error")
		return
	}
	_, err = ac.db.CreateOrgMember(r.Context(), sqlc.CreateOrgMemberParams{
		MemberID:  member.UserId,
		OrgID:     sqlOrg.OrgID,
		CreatorID: sqlOrg.UserID,
	})
	if err != nil {
		log.Print(err)
		utils.RespondWithError(w, http.StatusBadRequest, "Client error")
		return
	}

	payload := models.OrgMessageData{
		Status:  "success",
		Message: "User added to organisation successfully",
	}
	utils.RespondWithJSON(w, http.StatusOK, payload)
}
