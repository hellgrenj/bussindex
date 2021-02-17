package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hellgrenj/bussindex/pkg/errors"
	"github.com/hellgrenj/bussindex/pkg/system"
)

type operationResult struct {
	Result  interface{} `json:"result"`
	Success bool        `json:"success"`
}

func (s *Server) handleAPI(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&operationResult{Result: "API Alive and kicking"})
}

func (s *Server) createSystem(w http.ResponseWriter, r *http.Request) {

	var system system.System
	if err := s.decode(w, r, &system); err != nil {
		s.handleError(w, errors.NewInvalidError(err.Error(), err))
		return
	}
	id, err := s.systemService.Save(system)
	if err != nil {
		s.handleError(w, err)
		return
	}
	result := fmt.Sprintf("system created with id %v", id)
	s.respond(w, &operationResult{Result: result, Success: true})
}
func (s *Server) getSystems(w http.ResponseWriter, r *http.Request) {
	allSystems, err := s.systemService.Get()
	if err != nil {
		s.handleError(w, err)
	}
	s.respond(w, &operationResult{Result: allSystems, Success: true})
}
func (s *Server) respond(w http.ResponseWriter, response *operationResult) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func (s *Server) handleError(w http.ResponseWriter, err error) {
	if notFoundError, ok := err.(errors.NotFoundError); ok {
		s.error.Println(notFoundError.Error())
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(&operationResult{Result: notFoundError.ExternalError(), Success: false})
		return
	}
	if conflictError, ok := err.(errors.ConflictError); ok {
		s.error.Println(conflictError.Error())
		w.WriteHeader(409)
		json.NewEncoder(w).Encode(&operationResult{Result: conflictError.ExternalError(), Success: false})
		return
	}
	if invalidError, ok := err.(errors.InvalidError); ok {
		s.error.Println(invalidError.Error())
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&operationResult{Result: invalidError.ExternalError(), Success: false})
		return
	}

	w.WriteHeader(500)
	s.error.Println(err)
	json.NewEncoder(w).Encode(&operationResult{Result: "Internal Server Error", Success: false})
}
