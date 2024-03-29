package handlers

import (
	"encoding/json"
	"hellowWorldDeploy/pkg/entity"
	"net/http"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event := &entity.Event{}
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		h.l.Printf("error during decoding json in CreateEvent(handler): %v", err)
		h.WriteHTTPResponse(w, http.StatusBadRequest, "499")
		return
	}
	//log.Println(event)
	decodedClaims := r.Context().Value("decodedClaims").(*entity.Claims)
	event.CreatorID = decodedClaims.Sub
	if err := h.svc.CreateEvent(event); err != nil {
		h.l.Printf("error createEvent() CreateEvent(handler): %v", err)
		h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	//log.Println(event.ID, "----Handler")
	h.WriteHTTPResponse(w, http.StatusOK, "Event created!")
}

func (h *Handler) GetEventsByInterests(w http.ResponseWriter, r *http.Request) {
	var interests []string
	if err := json.NewDecoder(r.Body).Decode(&interests); err != nil {
		h.l.Printf("error during decoding json in GetEventsByInterests(handler): %v", err)
		h.WriteHTTPResponse(w, http.StatusBadRequest, "499")
		return
	}
	events, err := h.svc.GetEventsByInterests(interests)
	if err != nil {

		h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(events); err != nil {
		h.l.Printf("Error during sending response with %d: %v", http.StatusOK, err)
		return
	}
}
