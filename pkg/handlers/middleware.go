package handlers

import (
	"fmt"
	"hellowWorldDeploy/pkg/handlers/router"
	"net/http"
	"strings"
)

func (h *Handler) RunMiddlewares() router.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		route, err := h.router.GetRoute(r.URL.Path, r.Method)
		if err != nil {
			h.l.Print("NOT FOUND ------ from RunMiddlewares")
			http.NotFound(w, r)
			return
		}
		handler := route.Handler
		for _, middleware := range route.Middlewares {
			handler = middleware(handler)
		}
		handler(w, r)
		//fmt.Println("Successfully middleware stage")
		return
	}
}

func (h *Handler) AuthMiddleware(next router.Handler) router.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			h.l.Print("No Header Authorization in request ------ from AuthMiddleware")
			h.WriteHTTPResponse(w, http.StatusUnauthorized, "Authorization header is missing")
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			h.l.Print("No Bearer in token ------ from AuthMiddleware")
			h.WriteHTTPResponse(w, http.StatusUnauthorized, "Invalid authorization header format")
			return
		}
		fmt.Println(tokenStr, "-----This is TOKEEEEEEN")
		next(w, r)
	}
}
