package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kanatsanan6/todo-htmx-go/db"
	"github.com/kanatsanan6/todo-htmx-go/internal/repo"
)

type Server struct {
	router *chi.Mux
}

func NewServer() *Server {
	r := chi.NewRouter()

	return &Server{router: r}
}

func (s *Server) Start(port int, database *db.Database) error {
	r := s.router

	r.Use(middleware.Logger)

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	h := NewHealthHandler()
	r.Get("/health", h.Index())

	tr := repo.NewTaskRepo(database.GetDB())
	t := NewTaskHandler(tr)

	r.Get("/tasks", t.Index())
	r.Post("/tasks", t.Create())

	log.Printf("Listening to port %d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
