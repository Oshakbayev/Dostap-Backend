package pkg

import (
	"hellowWorldDeploy/pkg/middleware"
	"log"
	"net/http"
)

type Server struct {
	log        *log.Logger
	httpServer *http.Server
}

func InitServer(l *log.Logger) *Server {
	route := NewRouter()
	Routers(route)
	return &Server{
		log: l,
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: HTTPHandle(route),
		},
	}
}

func (s *Server) StartServer() error {
	log.Println("starting api server at http://localhost" + s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func HTTPHandle(route *Router) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", middleware.Middleware(route))
	return mux
}

func Routers(route *Router) {
	route.Get("/", MainPage)
}
