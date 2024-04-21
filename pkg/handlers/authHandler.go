package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"hellowWorldDeploy/pkg/entity"
	"net/http"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.l.Printf("error with decode data in SignUp(handler): %v", err)
		h.WriteHTTPResponse(w, http.StatusBadRequest, "499")
		return
	}
	if status, err := h.svc.SignUp(&user); err != nil {
		h.l.Printf("error with decode data in SignUp(handler): %v", err)
		h.WriteHTTPResponse(w, status, err.Error())
		return
	}
	h.WriteHTTPResponse(w, http.StatusOK, "verification email has been sent")
	return

}

func (h *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	var credentials entity.Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		h.WriteHTTPResponse(w, http.StatusBadRequest, "499")
		return
	}

	user, err := h.svc.LogIn(&credentials)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.l.Printf("ErrNoRows in LogIn(handler) %s", err.Error())
			h.WriteHTTPResponse(w, http.StatusBadRequest, "497")
			//return http.StatusBadRequest, entity.NilID, fmt.Errorf("497 no user exist with this email or username %s", credentials.Email+credentials.Username)
			return
		} else if err.Error() == "492" || err.Error() == "495" || err.Error() == "496" {
			h.WriteHTTPResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	signedToken, err := h.svc.TokenGenerator(user.ID, user.Email, user.Username)
	if err != nil {
		h.WriteHTTPResponse(w, http.StatusInternalServerError, err.Error())
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(struct {
		Token string
		Msg   string
	}{
		Token: signedToken,
		Msg:   "Your jwt token",
	}); err != nil {
		h.l.Printf("Error during sending response with %d: %v", http.StatusOK, err)
	}
}

func (h *Handler) TempHome(w http.ResponseWriter, r *http.Request) {
	h.WriteHTTPResponse(w, http.StatusOK, "YERKA pidr ne razreshil mainPage sdelat zashishennim((((")
}

func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) ConfirmAccount(w http.ResponseWriter, r *http.Request) {
	secretCode := r.FormValue("link")
	status, err := h.svc.VerifyAccount(secretCode)
	if err != nil {
		h.WriteHTTPResponse(w, status, err.Error())
		return
	}
	h.WriteHTTPResponse(w, http.StatusOK, "Your email is confirmed and you are registered!")
}
