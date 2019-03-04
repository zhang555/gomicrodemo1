package cerr

import "errors"

const (
//INTERNAL_SERVER_ERROR = "internal server error"
//INTERNAL_SERVER_ERROR = errors.New("internal server error")
)

var (
	ErrorInternalServerError = errors.New("internal server error")
)
