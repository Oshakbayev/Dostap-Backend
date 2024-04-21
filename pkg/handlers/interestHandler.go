package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"hellowWorldDeploy/pkg/entity"
	"net/http"
)

func (h *Handler) GetAllInterests(w http.ResponseWriter, r *http.Request) {
	allInterests, err := h.svc.GetAllInterests()
	if err != nil {
		h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err = json.NewEncoder(w).Encode(allInterests); err != nil {
		h.l.Printf("Error during sending response with %d: %v", http.StatusOK, err)
	}
}

func (h *Handler) GetUserInterests(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("decodedClaims").(*entity.Claims).Sub
	userIntersts, err := h.svc.GetUserInterests(userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			userIntersts = nil
		} else {
			h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	if err = json.NewEncoder(w).Encode(userIntersts); err != nil {
		h.l.Printf("Error during sending response with %d: %v", http.StatusOK, err)
	}
}
