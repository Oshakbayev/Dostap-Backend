package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var token map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		h.WriteHTTPResponse(w, http.StatusUnauthorized, err.Error())
		return
	}
	tokenStr := token["Token"].(string)
	_, status, err := h.svc.TokenChecker(tokenStr)
	if err != nil {
		h.WriteHTTPResponse(w, status, "Token "+tokenStr+" is invalid: "+err.Error())
		return
	}
	h.WriteHTTPResponse(w, http.StatusOK, "This is should be an array but I am too lazy and Huerka will not pay for this")
}
