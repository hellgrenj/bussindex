package system

import "github.com/hellgrenj/bussindex/pkg/validation"

// System is the stuct for the entity System
type System struct {
	ID          int64
	Description string
}

// OK is the validation function for the struct System
func (s *System) OK() error {
	if len(s.Description) == 0 {
		return validation.ErrMissingField("Description")
	}
	return nil
}
