package rest

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func (s *Server) routes() {
	s.router.Use(loggingMiddleware)
	s.router.HandleFunc("/api", s.handleAPI).Methods("GET")
	s.router.HandleFunc("/system", s.createSystem).Methods("POST")

}
func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, next)
}
