package pkg

import (
	configPkg "hellowWorldDeploy/cmd/config"
	"hellowWorldDeploy/database"
	"hellowWorldDeploy/pkg/handlers"
	"hellowWorldDeploy/pkg/handlers/router"
	repository "hellowWorldDeploy/pkg/repo"
	service2 "hellowWorldDeploy/pkg/service"
	"log"
	"net/http"
)

type Server struct {
	log        *log.Logger
	httpServer *http.Server
	handler    handlers.Handler
}

func InitServer(config *configPkg.Config, logger *log.Logger) *Server {
	db := database.CreateDB(config.DBDriver)
	repo := repository.CreateRepository(db, logger)
	service := service2.CreateService(repo, logger)
	route := router.NewRouter()
	handler := handlers.CreateHandler(service, route, logger)

	server := Server{
		log:     &log.Logger{},
		handler: handler,
	}

	handler.Routers()
	server.httpServer = &http.Server{
		Addr:    config.HTTPPort,
		Handler: handler.HTTPHandle(),
	}
	return &server
}

func (s *Server) StartServer() error {
	log.Println("starting api server at http://localhost" + s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
