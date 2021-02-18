package system

import "log"

// DBRepository is the interface for the system repository
type DBRepository interface {
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
	r    DBRepository
	info *log.Logger
}

// NewService creates a new system service
func NewService(r DBRepository, infoLogger *log.Logger) Service {
	return &service{r, infoLogger}
}
func (s *service) Save(system System) (int64, error) {
	s.info.Printf("saving a new system with description %s", system.Description)
	return s.r.Save(system)
}
func (s *service) Get() ([]System, error) {
	s.info.Println("fetching all systems")
	return s.r.Get()
}
func (s *service) Delete(id int) error {
	s.info.Printf("deleting system with id %d", id)
	return s.r.Delete(id)
}
