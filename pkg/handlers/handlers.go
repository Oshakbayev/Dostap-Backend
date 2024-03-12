package handlers

import (
	"encoding/json"
	"hellowWorldDeploy/pkg/entity"
	"hellowWorldDeploy/pkg/handlers/router"
	"hellowWorldDeploy/pkg/middleware"
	service "hellowWorldDeploy/pkg/service"
	"log"
	"net/http"
)

type Handler struct {
	svc   service.SvcInterface
	route *router.Router
	l     *log.Logger
}

func CreateHandler(svc service.SvcInterface, route *router.Router, l *log.Logger) Handler {
	return Handler{svc: svc, route: route, l: l}
}

func (h *Handler) HTTPHandle() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", middleware.Middleware(h.route))
	return mux
}

func (h *Handler) Routers() {
	h.route.Post("/signup", h.SignUp)
	h.route.Post("/login", h.LogIn)
	h.route.Get("/auth/confirmUserAccount", h.ConfirmAccount)
	h.route.Post("/updateProfile", h.ProfileEdit)
	h.route.Get("/createEvent", h.CreateEvent)
	h.route.Get("/", h.TempHome)
	//h.route.Post("/login", h.LogIn)

}

func (h *Handler) WriteHTTPResponse(w http.ResponseWriter, status int, msg string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(entity.ResponseJSON{Message: msg}); err != nil {
		h.l.Printf("Error during sending response with %d: %v", status, err)
	}
}
