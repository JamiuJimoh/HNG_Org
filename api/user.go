package api

import (
	"log"
	"net/http"

	"github.com/JamiuJimoh/hngorg/db/sqlc"
	"github.com/JamiuJimoh/hngorg/models"
	"github.com/JamiuJimoh/hngorg/utils"
)

func (ac *ApiCfg) GetUser(w http.ResponseWriter, r *http.Request) {
	reqUserId := r.PathValue("id")
	currentUserId := (ac.ctx.Value(currentUserIDKey)).(string)

	// get the organisation using any of the id if they are the same
	// if currentUserId == reqUserId {
	// 	ac.db.GetUserByID(r.Context(), reqUserId)
	// 	return
	// }
	// otherwise,
	// check if the reqUserId is in the organisation created by currentUserId
	// return the user details if true, respond with authorization error if false
	// return the user details if both IDs are in the same organisation
	sqlUser, err := ac.db.GetUserInSameOrgByID(r.Context(), sqlc.GetUserInSameOrgByIDParams{
		UserID:   currentUserId,
		UserID_2: reqUserId,
	})
	if err != nil {
		log.Print(currentUserId)
		utils.RespondWithError(w, http.StatusUnauthorized, "cannot fetch user not in the same organisation")
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
