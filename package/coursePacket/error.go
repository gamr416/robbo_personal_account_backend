package coursePacket

import "errors"

var (
	ErrInternalServerLevel  = errors.New("internal server error")
	ErrBadRequest           = errors.New("bad request")
	ErrBadRequestBody       = errors.New("bad request body")
	ErrCoursePacketNotFound = errors.New("course packet not found")
)
