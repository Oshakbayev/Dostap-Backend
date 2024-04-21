package handlers

import (
	"encoding/json"
	"hellowWorldDeploy/pkg/entity"
	"net/http"
)

func (h *Handler) ProfileEdit(w http.ResponseWriter, r *http.Request) {
	var updatedUser entity.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		h.WriteHTTPResponse(w, http.StatusBadRequest, "499")
		return
	}

	updatedUser.ID = r.Context().Value("decodedClaims").(*entity.Claims).Sub
	status, err := h.svc.UpdateUserProfileInfo(&updatedUser)
	if err != nil {
		h.WriteHTTPResponse(w, status, err.Error())
		return
	}
	h.WriteHTTPResponse(w, status, "success")
}

func (h *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	decodedClaims := r.Context().Value("decodedClaims").(*entity.Claims)
	err := h.svc.DeleteAccount(decodedClaims.Sub)
	if err != nil {
		h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.WriteHTTPResponse(w, http.StatusOK, "user deleted")
}
