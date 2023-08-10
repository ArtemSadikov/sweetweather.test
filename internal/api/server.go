package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sweetweather.test/internal/usecase"
)

type Server struct {
	uc usecase.UseCases
}

func (s *Server) Start(port int) error {
	router := mux.NewRouter()

	router.Use(authMiddleware)

	router.HandleFunc("/calculate", s.HandleCalculate()).Methods("POST")

	return http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func NewServer(uc usecase.UseCases) *Server {
	return &Server{uc}
}
