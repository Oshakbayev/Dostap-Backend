package pkg

import "net/http"

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

func (r *Router) getHandler(path, method string) {
	regexp := regexp.MustComipile
	for _, route := range r.routes {

	}
}
