package errset

import "fmt"

var (
	ErrUnAuthorized   = fmt.Errorf("unauthorized")
	ErrInternalServer = fmt.Errorf("internal server error")
	ErrBadRequest     = fmt.Errorf("bad request")
	ErrNotFound       = fmt.Errorf("not found")
)
