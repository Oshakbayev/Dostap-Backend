package handlers

import (
	"encoding/json"
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
