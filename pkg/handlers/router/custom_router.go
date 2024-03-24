package router

import (
	"errors"
	"net/http"
	"regexp"
)

type Handler func(w http.ResponseWriter, r *http.Request)

type Middleware func(handler Handler) Handler

type Route struct {
	Path        string
	Method      string
	Handler     Handler
	Middlewares []Middleware
}

type Router struct {
	routes []Route
}

func (r *Router) AddRoute(path string, method string, handler Handler, middlewares []Middleware) {
	r.routes = append(r.routes, Route{Path: path, Method: method, Handler: handler, Middlewares: middlewares})
}
func (r *Router) Get(path string, handler Handler, middlewares []Middleware) {
	r.AddRoute(path, http.MethodGet, handler, middlewares)
}

func (r *Router) Post(path string, handler Handler, middlewares []Middleware) {
	r.AddRoute(path, http.MethodPost, handler, middlewares)
}

func (r *Router) Put(path string, handler Handler, middlewares []Middleware) {
	r.AddRoute(path, http.MethodPut, handler, middlewares)
}

func (r *Router) GetRoute(path, method string) (Route, error) {

	for _, route := range r.routes {
		regex := regexp.MustCompile(route.Path)
		if regex.MatchString(path) && route.Method == method {
			return route, nil
		}

	}
	return Route{}, errors.New("route did not found")
}

//func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
//	path := req.URL.Path
//	method := req.Method
//	handler, err := r.getHandler(path, method)
//	//	fmt.Println(path, method, r)
//	if err != nil {
//		//	fmt.Println("not found")
//		http.NotFound(w, req)
//		return
//	}
//	handler(w, req)
//}

func NewRouter() *Router {
	return &Router{}
}
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}
