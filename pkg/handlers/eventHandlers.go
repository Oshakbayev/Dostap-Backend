package handlers

import (
	"encoding/json"
	"hellowWorldDeploy/pkg/entity"
	"log"
	"net/http"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event := &entity.Event{}
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		h.WriteHTTPResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Println(event)
	if err := h.svc.CreateEvent(event); err != nil {
		h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.WriteHTTPResponse(w, http.StatusOK, "This is should be an array but I am too lazy and Huerka will not pay for this")
}

func(h *Handler) GetEventsByInterests(w http.ResponseWriter, r *http.Request) {
	var interests []string
    if err := json.NewDecoder(r.Body).Decode(&interests); err!= nil {
        h.WriteHTTPResponse(w, http.StatusBadRequest, err.Error())
        return
    }
    events, err := h.svc.GetEventsByInterests(interests)
    if err!= nil {
        h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
        return
    }
    w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(struct {
		Events []entity.Event
	}{
		Events: events,
	}); err != nil {
		h.l.Printf("Error during sending response with %d: %v", http.StatusOK, err)
		return
	}
}
