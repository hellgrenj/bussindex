package system

import "github.com/hellgrenj/bussindex/pkg/validation"

// System is the stuct for the entity System
type System struct {
	ID     int64
	Name   string
	DevIds []int64
}

// OK is the validation function for the struct System
func (s *System) OK() error {
	if len(s.Name) == 0 {
		return validation.ErrMissingField("Name")
	}
	return nil
}
