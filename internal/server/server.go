package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/met4tron/doctors/config"
	"github.com/met4tron/doctors/internal/server/middlewares"
)

type Server struct {
	router *chi.Mux
	cfg    *config.Config
}

func newServer() *Server {
	return &Server{
		router: chi.NewRouter(),
		cfg:    config.LoadConfig(),
	}
}

func (s *Server) setMiddlewares() {
	middlewares.LoadMiddlewares(s.router)
}

func Init() *Server {
	server := newServer()
	server.setMiddlewares()

	return server
}

func (s *Server) Listen() {
	http.ListenAndServe(s.cfg.Port, s.router)
}
