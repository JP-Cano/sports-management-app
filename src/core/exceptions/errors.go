package exceptions

import (
	"errors"
)

type Error struct {
	Err error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func (e *Error) Is(target error) bool {
	var t *Error
	ok := errors.As(target, &t)
	if !ok {
		return false
	}
	return errors.Is(e.Err, t.Err)
}

var (
	InternalServerError = &Error{Err: errors.New("internal Server Error")}
	NotFound            = &Error{Err: errors.New("not Found")}
	BadRequest          = &Error{Err: errors.New("bad request")}
)

func Throw(err *Error) *Error {
	return &Error{Err: err.Err}
}
