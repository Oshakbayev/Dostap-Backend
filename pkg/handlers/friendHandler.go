package handlers

import (
	"encoding/json"
	"hellowWorldDeploy/pkg/entity"
	"net/http"
)

func (h *Handler) CreateFriendRequest(w http.ResponseWriter, r *http.Request) {
	var fRequest entity.FriendRequest
	if err := json.NewDecoder(r.Body).Decode(&fRequest); err != nil {
		h.WriteHTTPResponse(w, http.StatusBadRequest, "499")
		return
	}
	fRequest.SenderID = r.Context().Value("decodedClaims").(*entity.Claims).Sub
	if err := h.svc.CreateFriendRequest(fRequest); err != nil {
		h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.WriteHTTPResponse(w, http.StatusOK, "friend request created!")
}

func (h *Handler) FriendRequestAnswer(w http.ResponseWriter, r *http.Request) {
	var reqBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		h.WriteHTTPResponse(w, http.StatusBadRequest, "499")
		return
	}
	temp := reqBody["fRequestID"].(float64)
	fRequestID := int64(temp)
	fRequestAnswer := reqBody["answer"].(bool)
	if err := h.svc.EditFriendRequestStatus(fRequestID, fRequestAnswer); err != nil {
		h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.WriteHTTPResponse(w, http.StatusOK, "friends created!")
}

func (h *Handler) GetFriendRequests(w http.ResponseWriter, r *http.Request) {
	recipientID := r.Context().Value("decodedClaims").(*entity.Claims).Sub
	friendRequests, err := h.svc.GetFriendRequestsByRecipientID(recipientID)
	if err != nil {
		h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err = json.NewEncoder(w).Encode(friendRequests); err != nil {
		h.l.Printf("Error during sending response with %d: %v", http.StatusOK, err)
		return
	}
}

func (h *Handler) DeleteFriend(w http.ResponseWriter, r *http.Request) {
	friendID1 := r.Context().Value("decodedClaims").(*entity.Claims).Sub
	var reqBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		h.WriteHTTPResponse(w, http.StatusBadRequest, "499")
		return
	}
	friendID2 := int64(reqBody["friendID"].(float64))
	err := h.svc.DeleteFriend(friendID1, friendID2)
	if err != nil {
		h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.WriteHTTPResponse(w, http.StatusOK, "friend deleted!")
}
