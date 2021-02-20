package system

import (
	"log"
	"strings"
)

// IRepository is the interface for the system repository
type IRepository interface {
	Save(System) (int64, error)
	Get() ([]System, error)
	Delete(id int) error
}

// Service provides system features
type Service interface {
	Save(System) (int64, error)
	Get() ([]System, error)
	Delete(id int) error
}
type service struct {
	r    IRepository
	info *log.Logger
}

// NewService creates a new system service
func NewService(r IRepository, infoLogger *log.Logger) Service {
	return &service{r, infoLogger}
}
func (s *service) Save(system System) (int64, error) {
	s.info.Printf("saving a new system with name %s", system.Name)
	return s.r.Save(system)
}
func (s *service) Get() ([]System, error) {
	s.info.Println("fetching all systems")
	allSystems, err := s.r.Get()
	if err != nil {
		return nil, err
	}
	for index := range allSystems {
		allSystems[index].Name = strings.ToLower(allSystems[index].Name)
	}
	return allSystems, nil
}
func (s *service) Delete(id int) error {
	s.info.Printf("deleting system with id %d", id)
	return s.r.Delete(id)
}
