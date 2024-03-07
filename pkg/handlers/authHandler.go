package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"hellowWorldDeploy/pkg/entity"
	"net/http"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {

	user := entity.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		//fmt.Println("bad request is correct")
		return
	}
	fmt.Println(user)
	if err := h.svc.SignUp(&user); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user uspeshno zaregan"))
	return

}

func (h *Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	var credentials entity.Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	phoneNum := credentials.PhoneNum
	pass := credentials.Password
	signedToken, err := h.svc.LogIn(phoneNum, pass)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Creation of jwt token failed"))
		fmt.Println("HERE?")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(struct {
		Token string
	}{
		Token: signedToken,
	}); err != nil {
		fmt.Println("TOKEEEEEN DID NOT SENT", err)
	}
}

func (h *Handler) TempHome(w http.ResponseWriter, r *http.Request) {
	var token map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	tokenStr := token["Token"].(string)
	claims := &entity.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return entity.JWTKey, nil
	})
	if err != nil {
		fmt.Println("{", err, "}", "{", jwt.ErrSignatureInvalid, "}")
		if err.Error() == jwt.ErrSignatureInvalid.Error() {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprintf("bad signature")))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("bad request")))
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprintf("invalid token")))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello %s", claims.PhoneNum)))
}

func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {

}
