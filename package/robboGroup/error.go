package robboGroup

import "errors"

var (
	ErrInternalServerLevel = errors.New("internal server level")
	ErrBadRequest          = errors.New("bad request")
	ErrBadRequestBody      = errors.New("bad request body")
	ErrRobboGroupNotFound  = errors.New("robbo group not found")
)
