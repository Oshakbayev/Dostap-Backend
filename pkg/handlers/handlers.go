package handlers

import (
	"hellowWorldDeploy/pkg/handlers/router"
	"hellowWorldDeploy/pkg/middleware"
	service "hellowWorldDeploy/pkg/service"
	"net/http"
)

type Handler struct {
	svc   service.SvcInterface
	route *router.Router
}

func CreateHandler(svc service.SvcInterface, route *router.Router) Handler {
	return Handler{svc: svc, route: route}
}

func (h *Handler) HTTPHandle() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", middleware.Middleware(h.route))
	return mux
}

func (h *Handler) Routers() {
	h.route.Post("/signup", h.SignUp)
	h.route.Post("/login", h.LogIn)
	//h.route.Post("/login", h.LogIn)

}
