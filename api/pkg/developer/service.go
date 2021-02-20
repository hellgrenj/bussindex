package developer

import (
	"log"
)

// IRepository is the interface for the system repository
type IRepository interface {
	Save(Developer) (int64, error)
	Get() ([]Developer, error)
	Delete(id int) error
}

// Service provides system features
type Service interface {
	Save(Developer) (int64, error)
	Get() ([]Developer, error)
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
func (s *service) Save(developer Developer) (int64, error) {
	s.info.Printf("saving a new developer with name %s", developer.Name)
	return s.r.Save(developer)
}
func (s *service) Get() ([]Developer, error) {
	s.info.Println("fetching all developers")
	return s.r.Get()
}
func (s *service) Delete(id int) error {
	s.info.Printf("deleting developer with id %d", id)
	return s.r.Delete(id)
}
