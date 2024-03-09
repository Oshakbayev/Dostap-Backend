package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"hellowWorldDeploy/pkg/entity"
	"net/http"
)

func (h *Handler) ProfileEdit(w http.ResponseWriter, r *http.Request) {
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
	user := entity.User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		//fmt.Println("bad request is correct")
		return
	}
	status, err := h.svc.UpdateUserProfileInfo(&user)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(status)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte("Information Updates succesfully"))
}
