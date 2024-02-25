package pkg

import (
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func InitServer() *Server {
	return &Server{httpServer: &http.Server{
		Addr: "80",
	},
	}
}

func (s *Server) StartServer() error {
	log.Println("starting api server at http://localhost" + s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) InitRouter() error {
	route := NewRouter()
	s.Routers(route)
	return s.StartServer()

}

func (s *Server) Routers(route *Router) {
	route.Get("/", MainPage)
}
