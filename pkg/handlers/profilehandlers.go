package handlers

import (
	"encoding/json"
	"hellowWorldDeploy/pkg/entity"
	"net/http"
)

func (h *Handler) ProfileEdit(w http.ResponseWriter, r *http.Request) {
	var updateJson entity.UpdateJson
	if err := json.NewDecoder(r.Body).Decode(&updateJson); err != nil {
		h.WriteHTTPResponse(w, http.StatusUnauthorized, err.Error())
		return
	}
	user := updateJson.UserInfo
	decodedClaims, status, err := h.svc.TokenChecker(updateJson.Token.Token)
	if err != nil {
		h.WriteHTTPResponse(w, status, "Token "+updateJson.Token.Token+" is invalid: "+err.Error())
		return
	}
	user.ID = decodedClaims.Sub
	status, err = h.svc.UpdateUserProfileInfo(&user)
	if err != nil {
		h.WriteHTTPResponse(w, status, err.Error())
		return
	}
	h.WriteHTTPResponse(w, status, "success")
}

func (h *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	var updateJson entity.UpdateJson
	if err := json.NewDecoder(r.Body).Decode(&updateJson); err != nil {
		h.WriteHTTPResponse(w, http.StatusUnauthorized, err.Error())
		return
	}
	decodedClaims, status, err := h.svc.TokenChecker(updateJson.Token.Token)
	if err != nil {
		h.WriteHTTPResponse(w, status, "Token "+updateJson.Token.Token+" is invalid: "+err.Error())
		return
	}
	err = h.svc.DeleteAccount(decodedClaims.Sub)
	if err != nil {
		h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.WriteHTTPResponse(w, http.StatusOK, "user deleted")
}
