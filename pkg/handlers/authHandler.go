package handlers

import (
	"encoding/json"
	"fmt"
	"hellowWorldDeploy/pkg/entity"
	"net/http"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.WriteHTTPResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if status, err := h.svc.SignUp(&user); err != nil {
		h.WriteHTTPResponse(w, status, err.Error())
		return
	}
	h.WriteHTTPResponse(w, http.StatusOK, "verification email has been sent")
	return

}

func (h *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	var credentials entity.Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		h.WriteHTTPResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	email := credentials.Email
	pass := credentials.Password
	status, userID, err := h.svc.LogIn(email, pass)
	if err != nil {
		h.WriteHTTPResponse(w, status, err.Error())
		return
	}
	signedToken, err := h.svc.TokenGenerator(userID, email)
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
		fmt.Println("token did not send", err)
	}
}

func (h *Handler) TempHome(w http.ResponseWriter, r *http.Request) {
	h.WriteHTTPResponse(w, http.StatusOK, "YERKA pidr ne razreshil mainPage sdelat zashishennim((((")
}

func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) ConfirmAccount(w http.ResponseWriter, r *http.Request) {
	secretCode := r.FormValue("link")
	fmt.Println(secretCode)
	status, err := h.svc.VerifyAccount(secretCode)
	if err != nil {
		h.WriteHTTPResponse(w, status, err.Error())
		return
	}
	h.WriteHTTPResponse(w, http.StatusOK, "Your email is confirmed and you are registered!")
}
