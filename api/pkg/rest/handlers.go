package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hellgrenj/bussindex/pkg/developer"
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
func (s *Server) createDeveloper(w http.ResponseWriter, r *http.Request) {
	var developer developer.Developer
	if err := s.decode(w, r, &developer); err != nil {
		s.handleError(w, errors.NewInvalidError(err.Error(), err))
		return
	}
	id, err := s.developerService.Save(developer)
	if err != nil {
		s.handleError(w, err)
		return
	}
	result := fmt.Sprintf("developer created with id %v", id)
	s.respond(w, &operationResult{Result: result, Success: true})
}
func (s *Server) getSystems(w http.ResponseWriter, r *http.Request) {
	allSystems, err := s.systemService.Get()
	if err != nil {
		s.handleError(w, err)
		return
	}
	s.respond(w, &operationResult{Result: allSystems, Success: true})
}
func (s *Server) getDevelopers(w http.ResponseWriter, r *http.Request) {
	allDevelopers, err := s.developerService.Get()
	if err != nil {
		s.handleError(w, err)
		return
	}
	s.respond(w, &operationResult{Result: allDevelopers, Success: true})
}
func (s *Server) deleteSystemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.handleError(w, errors.NewInvalidError("You need to pass in an id", err))
		return
	}
	err = s.systemService.Delete(id)
	if err != nil {
		s.handleError(w, err)
		return
	}
	s.respond(w, &operationResult{Result: "system deleted", Success: true})
}
func (s *Server) addDeveloperToSystem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	systemID, err := strconv.Atoi(vars["systemID"])
	if err != nil {
		s.handleError(w, errors.NewInvalidError("You need to pass in an systemID", err))
		return
	}
	developerID, err := strconv.Atoi(vars["developerID"])
	if err != nil {
		s.handleError(w, errors.NewInvalidError("You need to pass in an developerID", err))
		return
	}
	err = s.systemService.AddDeveloper(systemID, developerID)
	if err != nil {
		s.handleError(w, err)
		return
	}
	s.respond(w, &operationResult{Result: "developer added to system", Success: true})
}
func (s *Server) removeDeveloperFromSystem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	systemID, err := strconv.Atoi(vars["systemID"])
	if err != nil {
		s.handleError(w, errors.NewInvalidError("You need to pass in an systemID", err))
		return
	}
	developerID, err := strconv.Atoi(vars["developerID"])
	if err != nil {
		s.handleError(w, errors.NewInvalidError("You need to pass in an developerID", err))
		return
	}
	err = s.systemService.RemoveDeveloper(systemID, developerID)
	if err != nil {
		s.handleError(w, err)
		return
	}
	s.respond(w, &operationResult{Result: "developer removed from system", Success: true})
}
func (s *Server) getDevIdsWorkingOnSystem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	systemID, err := strconv.Atoi(vars["systemID"])
	if err != nil {
		s.handleError(w, errors.NewInvalidError("You need to pass in an systemID", err))
		return
	}

	devIds, err := s.systemService.GetDevIdsWorkingOnSystem(systemID)
	if err != nil {
		s.handleError(w, err)
		return
	}
	s.respond(w, &operationResult{Result: devIds, Success: true})
}
func (s *Server) deleteDeveloperByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		s.handleError(w, errors.NewInvalidError("You need to pass in an id", err))
	}
	err = s.developerService.Delete(id)
	if err != nil {
		s.handleError(w, err)
	}
	s.respond(w, &operationResult{Result: "developer deleted", Success: true})
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
