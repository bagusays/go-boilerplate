package errors

import "errors"

var (
	ErrOthers = errors.New("internal server error")
)

var mapErrToStatusCode = map[error]int{
	ErrOthers: 99,
}

func MapErrToStatusCode(err error) int {
	statusCode := mapErrToStatusCode[err]
	return statusCode
}
