package handlers

import (
	"encoding/json"
	"hellowWorldDeploy/pkg/entity"
	"hellowWorldDeploy/pkg/handlers/router"
	service "hellowWorldDeploy/pkg/service"
	"log"
	"net/http"
)

type Handler struct {
	svc    service.SvcInterface
	router *router.Router
	l      *log.Logger
}

func CreateHandler(svc service.SvcInterface, route *router.Router, l *log.Logger) Handler {
	return Handler{svc: svc, router: route, l: l}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", h.RunMiddlewares())
	return mux
}

func (h *Handler) Routers() {
	h.router.Post("/signup", h.SignUp, nil)
	h.router.Post("/login", h.LogIn, nil)
	h.router.Get("/auth/confirmUserAccount", h.ConfirmAccount, nil)
	h.router.Put("/updateProfile", h.ProfileEdit, []router.Middleware{h.AuthMiddleware})
	h.router.Post("/createEvent", h.CreateEvent, []router.Middleware{h.AuthMiddleware})
	h.router.Put("/deleteAccount", h.DeleteAccount, []router.Middleware{h.AuthMiddleware})
	h.router.Post("/getEventsByInterests", h.GetEventsByInterests, []router.Middleware{h.AuthMiddleware})
	h.router.Post("/createFriendRequest", h.CreateFriendRequest, []router.Middleware{h.AuthMiddleware})
	h.router.Put("/friendRequestAnswer", h.FriendRequestAnswer, []router.Middleware{h.AuthMiddleware})
	h.router.Get("/getFriendRequests", h.GetFriendRequests, []router.Middleware{h.AuthMiddleware})
	h.router.Post("/deleteFriend", h.DeleteFriend, []router.Middleware{h.AuthMiddleware})
	h.router.Get("/getAllInterests", h.GetAllInterests, []router.Middleware{h.AuthMiddleware})
	h.router.Get("/getAllEvents", h.GetAllEvents, []router.Middleware{h.AuthMiddleware})
	h.router.Get("/getAllUsernames", h.GetAllUsernames, nil)
	h.router.Post("/uploadFile", h.UploadFile, nil)
	h.router.Get("/getEventsByPage", h.GetEventsByPage, []router.Middleware{h.AuthMiddleware})
	//h.route.Post("/login", h.LogIn)
	//h.router.Get("/", h.TempHome, nil)

}

//example of json body for CreateEvent function

func (h *Handler) WriteHTTPResponse(w http.ResponseWriter, status int, msg string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(entity.ResponseJSON{Message: msg}); err != nil {
		h.l.Printf("Error during sending response with %d: %v", status, err)
	}
}
