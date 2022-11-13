package courses

import "errors"

var (
	ErrBadRequest     = errors.New("bad request")
	ErrInternalServer = errors.New("internal server error")
	ErrBadRequestBody = errors.New("bad request body")
)
