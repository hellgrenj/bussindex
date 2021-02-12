package errors

import "fmt"

// InvalidError is a domain error for when there is an invalid request
type InvalidError struct {
	expl string
	err  error
}

func (i InvalidError) Error() string {
	return fmt.Sprintf("%s, %v", i.expl, i.err)
}

// ExternalError prints an error safe for external use (in an API response for example)
func (i InvalidError) ExternalError() string {
	return fmt.Sprintf("%s", i.expl)
}

//NewInvalidError constructs a new InvalidError with an explanation
func NewInvalidError(expl string, e error) InvalidError {
	return InvalidError{expl, e}
}
