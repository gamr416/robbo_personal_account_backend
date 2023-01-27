package courses

import "errors"

var (
	ErrBadRequest          = errors.New("bad request")
	ErrInternalServerLevel = errors.New("internal server error")
	ErrBadRequestBody      = errors.New("bad request body")
	ErrCourseNotFound      = errors.New("course not found")
	ErrIncorrectInputParam = errors.New("error incorrect input params")
)
