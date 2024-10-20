package server

import (
	"fmt"
	"net/http"

	"github.com/sagoresarker/svc-discovery-host-poc/internal/config"
	"github.com/sagoresarker/svc-discovery-host-poc/internal/discovery"
)

type Server struct {
	cfg     *config.Config
	handler *discovery.Handler
	service *discovery.Service
}

func NewServer(cfg *config.Config) *Server {
	service := discovery.NewService()
	handler := discovery.NewHandler(service)

	return &Server{
		cfg:     cfg,
		handler: handler,
		service: service,
	}
}

func (s *Server) Run() error {
	http.HandleFunc("/register", s.handler.RegisterAgent)
	http.HandleFunc("/agents", s.handler.GetAllAgents)
	http.HandleFunc("/health", s.healthHandler)

	fmt.Printf("Server is running on port %d\n", s.cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.Port), nil)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Host service is healthy"))
}
