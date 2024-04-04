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
	event.OrganizerIDs = append(event.OrganizerIDs, decodedClaims.Username)
	//log.Println(decodedClaims.Sub, decodedClaims.Username)
	if err := h.svc.CreateEvent(event); err != nil {
		h.l.Printf("error createEvent() CreateEvent(handler): %v", err)
		h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	//log.Println(event.ID, "----Handler")
	h.WriteHTTPResponse(w, http.StatusOK, "Event created!")
}


func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.l.Printf("error during parsing form in UploadFile(handler): %v", err)
        h.WriteHTTPResponse(w, http.StatusBadRequest, "499")
        return
	}

	file, _, err := r.FormFile("picture")
	if err!= nil {
        h.l.Printf("error during reading file in UploadFile(handler): %v", err)
        h.WriteHTTPResponse(w, http.StatusBadRequest, "499")
        return
    }
	defer file.Close()
	h.svc.UploadFile(file)

	

	h.WriteHTTPResponse(w, http.StatusOK, "File uploaded!")

	
}

func (h *Handler) GetEventsByInterests(w http.ResponseWriter, r *http.Request) {
	var interests []int
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

func (h *Handler) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.svc.GetAllEvents()
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
