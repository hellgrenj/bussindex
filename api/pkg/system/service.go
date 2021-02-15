package system

import "log"

// DBRepository is the interface for the system repository
type DBRepository interface {
	Save(System) (int64, error)
}

// Service provides system features
type Service interface {
	Save(System) (int64, error)
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
	s.info.Println("SAD PANDA4")
	return s.r.Save(system)
}
