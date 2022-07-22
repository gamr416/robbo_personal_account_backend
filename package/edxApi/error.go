package edxApi

import "errors"

var (
	ErrIncorrectInputParam = errors.New("error incorrect input params")
	ErrOnReq               = errors.New("error on request")
	ErrReadRespBody        = errors.New("error while reading the response bytes")
	ErrTknNotRefresh       = errors.New("token not refresh")
	ErrOnResp              = errors.New("error on response")
	ErrJsonMarshal         = errors.New("error on json Marshal")
)
