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
