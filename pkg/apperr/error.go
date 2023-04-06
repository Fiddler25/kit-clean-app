package apperr

import "errors"

var (
	ErrBadRoute        = errors.New("bad route")
	ErrInvalidArgument = errors.New("invalid argument")
)
