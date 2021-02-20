package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/golang/gddo/httputil/header"
	"github.com/gorilla/mux"
	"github.com/hellgrenj/bussindex/pkg/developer"
	"github.com/hellgrenj/bussindex/pkg/system"
	"github.com/hellgrenj/bussindex/pkg/validation"
)

// Server is the http server struct
type Server struct {
	router           *mux.Router
	systemService    system.Service
	developerService developer.Service
	info             *log.Logger
	error            *log.Logger
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	s.router.ServeHTTP(w, r)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS, PUT, PATCH")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Access-Control-Allow-Methods, Authorization, X-Requested-With")
}

func (s *Server) decode(w http.ResponseWriter, r *http.Request, v validation.Ok) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			return errors.New(msg)
		}
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576) // request body max 1 MB
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(v); err != nil {
		return err
	}
	return v.OK()
}

// NewServer returns a new http server
func NewServer(systemService system.Service, developerService developer.Service, infoLogger *log.Logger, errorLogger *log.Logger) *Server {
	s := &Server{router: mux.NewRouter(), systemService: systemService, developerService: developerService, info: infoLogger, error: errorLogger}
	s.routes()
	return s
}

// StartAndListen starts the server and listens on the provided port
func (s *Server) StartAndListen(port int) {
	s.info.Println(fmt.Sprintf("API running on port %d", port))
	s.error.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), s))
}
