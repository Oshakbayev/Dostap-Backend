package handlers

import (
	"encoding/json"
	"fmt"
	"hellowWorldDeploy/pkg/entity"
	"net/http"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	//authotorization check
	//authHeader := r.Header.Get("Authorization")
	//if authHeader == "" {
	//	h.WriteHTTPResponse(w, http.StatusUnauthorized, "Authorization header is missing")
	//	return
	//}
	//tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	//if tokenStr == authHeader {
	//	h.WriteHTTPResponse(w, http.StatusUnauthorized, "Invalid authorization header format")
	//	return
	//}
	//fmt.Println(tokenStr, "This is TOKEEEEEEN")
	//_, status, err := h.svc.TokenChecker(tokenStr)
	//if err != nil {
	//	h.WriteHTTPResponse(w, status, "Token "+tokenStr+" is invalid: "+err.Error())
	//	return
	//}

	event := &entity.Event{}
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		h.WriteHTTPResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(event)
	if err := h.svc.CreateEvent(event); err != nil {
		h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.WriteHTTPResponse(w, http.StatusOK, "This is should be an array but I am too lazy and Huerka will not pay for this")
}
