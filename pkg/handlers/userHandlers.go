package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetAllUsernames(w http.ResponseWriter, r *http.Request) {
	userNames, err := h.svc.GetAllUsernames()
	if err!= nil {
        h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
        return
    }
	if err = json.NewEncoder(w).Encode(userNames); err!= nil {
        h.l.Printf("Error during sending response with %d: %v", http.StatusOK, err)
    }
}