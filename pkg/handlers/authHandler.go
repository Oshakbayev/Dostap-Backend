package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"hellowWorldDeploy/pkg/entity"
	"net/http"
	"time"
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
	if err := h.svc.LogIn(phoneNum, pass); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		fmt.Println("HERE?")
		return
	}
	expTime := time.Now().Add(time.Minute * 5)
	claims := &entity.Claims{
		PhoneNum: phoneNum,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(entity.JWTKey)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Println(signedToken)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(struct {
		Token string
	}{
		Token: signedToken,
	}); err != nil {
		fmt.Println("TOKEEEEEN DID NOT SENT", err)
	}
}
