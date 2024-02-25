package pkg

import (
	"errors"
	"net/http"
	"regexp"
)

type Handler func(w http.ResponseWriter, r *http.Request)

type route struct {
	Path    string
	Method  string
	Handler Handler
}

type Router struct {
	routes []route
}

func (r *Router) AddRoute(path string, method string, handler Handler) {
	r.routes = append(r.routes, route{Path: path, Method: method, Handler: handler})
}
func (r *Router) Get(path string, handler Handler) {
	r.AddRoute(path, http.MethodGet, handler)
}

func (r *Router) Post(path string, handler Handler) {
	r.AddRoute(path, http.MethodPost, handler)
}

func (r *Router) getHandler(path, method string) (Handler, error) {
	regex := regexp.MustCompile(path)
	for _, route := range r.routes {
		if regex.MatchString(route.Path) && route.Method == method {
		}
		return route.Handler, nil
	}
	return nil, errors.New("Handler did not found")
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method
	handler, err := r.getHandler(path, method)
	if err != nil {
		http.NotFound(w, req)
		return
	}
	handler(w, req)
}

func NewRouter() *Router {
	return &Router{}
}
