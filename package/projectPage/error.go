package projectPage

import "errors"

var (
	ErrPageNotFound        = errors.New("page not found")
	ErrInternalServerLevel = errors.New("internal server level")
	ErrBadRequest          = errors.New("bad request")
)
