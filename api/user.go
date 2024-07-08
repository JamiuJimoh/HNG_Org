package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JamiuJimoh/hngorg/db/sqlc"
	"github.com/JamiuJimoh/hngorg/models"
	"github.com/JamiuJimoh/hngorg/utils"
)

func (ac *ApiCfg) GetUserInSameOrg(w http.ResponseWriter, r *http.Request) {
	reqUserId := r.PathValue("id")
	currentUserId := (ac.ctx.Value(currentUserIDKey)).(string)

	sqlUser, err := ac.db.GetUserInSameOrgByID(r.Context(), sqlc.GetUserInSameOrgByIDParams{
		MemberID:   currentUserId,
		MemberID_2: reqUserId,
	})
	if err != nil {
		log.Print(err)
		utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("Found no user with id: %v in the same organisation", reqUserId))
		return
	}

	user := models.ResUserFromSQLUserSameOrgRow(sqlUser)
	payload := &models.SingleUserResData{
		Status:  "success",
		Message: "found user",
		Data:    user,
	}
	utils.RespondWithJSON(w, http.StatusOK, payload)
}
