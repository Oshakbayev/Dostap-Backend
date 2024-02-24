//package pkg
//
//import (
//	"log"
//	"net/http"
//)
//
//type Server struct {
//	httpServer *http.Server
//}
//
//func InitServer() *Server {
//	return &Server{httpServer: &http.Server{
//		Addr:    "8080",
//		Handler: InitRouter(),
//	},
//	}
//}
//
//func (s *Server) StartServer() error {
//	log.Println("starting api server at http://localhost" + s.httpServer.Addr)
//	return s.httpServer.ListenAndServe()
//}
//
////func InitRouter() *http.ServeMux {
////	mux := http.NewServeMux()
////
////}
