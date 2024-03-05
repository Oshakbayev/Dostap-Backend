package middleware

import (
	"net/http"
)

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//middlware for cokkie
		handler.ServeHTTP(w, r)
		//fmt.Println("Successfully middleware stage")
		return
	})
}
