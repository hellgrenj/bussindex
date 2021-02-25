package developer

import (
	"time"

	"github.com/hellgrenj/bussindex/pkg/validation"
)

// Developer is the struct for the entity developer
type Developer struct {
	ID               int64
	Name             string
	DateOfEmployment time.Time
}

// OK is the validation function for the struct System
func (d *Developer) OK() error {
	if len(d.Name) == 0 {
		return validation.ErrMissingField("Name")
	}
	if d.DateOfEmployment.IsZero() {
		return validation.ErrMissingField("DateOfEmployment")
	}
	return nil
}
