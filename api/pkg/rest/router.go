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
	s.router.HandleFunc("/system/{id}", s.deleteSystemByID).Methods("DELETE")
	s.router.HandleFunc("/system", s.getSystems).Methods("GET")
	s.router.HandleFunc("/system/{systemID}/adddev/{developerID}", s.addDeveloperToSystem).Methods("PUT")
	s.router.HandleFunc("/system/{systemID}/removedev/{developerID}", s.removeDeveloperFromSystem).Methods("PUT")
	s.router.HandleFunc("/system/{systemID}/devs", s.getDevIdsWorkingOnSystem).Methods("GET")

	s.router.HandleFunc("/developer", s.createDeveloper).Methods("POST")
	s.router.HandleFunc("/developer", s.getDevelopers).Methods("GET")
	s.router.HandleFunc("/developer/{id}", s.deleteDeveloperByID).Methods("DELETE")

}
func loggingMiddleware(next http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, next)
}
