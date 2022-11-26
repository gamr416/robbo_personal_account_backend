package robboUnits

import "errors"

var (
	ErrInternalServerLevel = errors.New("internal server level")
	ErrBadRequest          = errors.New("bad request")
	ErrBadRequestBody      = errors.New("bad request body")
	ErrRobboUnitNotFound   = errors.New("robbo unit not found")
)
