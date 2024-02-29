package middleware

import (
	"fmt"
	"net/http"
)

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		fmt.Println("Successfully middleware stage")
		return
	})
}
