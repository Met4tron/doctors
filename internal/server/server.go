package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/met4tron/doctors/config"
	"github.com/met4tron/doctors/internal/server/middlewares"
	"github.com/met4tron/doctors/pkg/logger"
	"go.uber.org/zap"
)

type Server struct {
	router *chi.Mux
	cfg    *config.Config
	logger *zap.SugaredLogger
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

func (s *Server) runMigrations() {
	urlString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		s.cfg.DatabaseConfig.User, s.cfg.DatabaseConfig.Password, s.cfg.DatabaseConfig.Host, s.cfg.DatabaseConfig.Port, s.cfg.DatabaseConfig.Schema)

	db, err := sql.Open("postgres", urlString)

	if err != nil {
		log.Panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/db/migrations",
		"postgres", driver)

	if err != nil {
		log.Println("Error rolou")
		log.Println(err)
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	defer db.Close()
}

func Init() *Server {
	logger.InitLogger()

	server := newServer()
	server.setMiddlewares()
	server.runMigrations()

	return server
}

func (s *Server) Listen() {
	logger.Info("Server initialising")
	logger.Fatal("Server closed", http.ListenAndServe(s.cfg.ApiConfig.Port, s.router))
}
