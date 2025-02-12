package server

import (
	"net/http"
	"strings"

	"go-weather-webapp/internal/handlers"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

func (s *Server) SetupRoutes() {
	s.router.HandleFunc("/weather/{zipcode}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		zipcode := strings.TrimSpace(vars["zipcode"])
		handlers.WeatherHandler(w, r, zipcode)
	}).Methods("GET")
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}

func (s *Server) Stop() {
	// Logic to gracefully stop the server can be implemented here
}
