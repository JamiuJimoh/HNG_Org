package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/JamiuJimoh/hngorg/db/sqlc"
	"github.com/JamiuJimoh/hngorg/models"
	"github.com/JamiuJimoh/hngorg/utils"
	"github.com/jackc/pgx/v5/pgtype"
)

func (ac *ApiCfg) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	read, err := io.ReadAll(r.Body)
	if err != nil {
		handleRegistrationError(w, err)
		return
	}
	json.Unmarshal(read, &user)

	w.Header().Set("Content-Type", "application/json")
	validatorError := utils.ValidateUserModel(user)
	if validatorError != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(validatorError)
		return
	}
	err = user.SanitizeUser()
	if err != nil {
		handleRegistrationError(w, err)
		return
	}

	org, err := user.CreateOrgFromUser()
	if err != nil {
		handleRegistrationError(w, err)
		return
	}

	dbUser, err := ac.db.CreateUser(r.Context(), sqlc.CreateUserParams{
		ID:        user.UserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     pgtype.Text{String: user.Phone, Valid: true},
	})
	if err != nil {
		handleRegistrationError(w, err)
		return
	}

	_, err = ac.db.CreateOrganisation(r.Context(), sqlc.CreateOrganisationParams{
		OrgID:       org.OrgId,
		Name:        org.Name,
		Description: pgtype.Text{String: org.Description, Valid: true},
		UserID:      dbUser.ID,
	})

	orgMember := user.CreateOrgMemberFromUser(org.OrgId, user.UserId)
	_, err = ac.db.CreateOrgMember(r.Context(), sqlc.CreateOrgMemberParams{
		MemberID:  orgMember.Id,
		OrgID:     orgMember.OrgId,
		CreatorID: orgMember.Creator_id,
	})
	if err != nil {
		handleRegistrationError(w, err)
		return
	}

	payload, err := ac.createLoginData(models.ResUserFromDBUser(dbUser), "Registration successful")
	if err != nil {
		handleRegistrationError(w, err)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, payload)
}

func handleRegistrationError(w http.ResponseWriter, err error) {
	log.Print(err)
	utils.RespondWithError(w, http.StatusBadRequest, "Registration unsuccessful")
	// http.Error(w, "An error occured", http.StatusInternalServerError)
}
