package errors

import "errors"

var (
	ErrPathParameter     = errors.New("path parameter error")
	ErrNotFound          = errors.New("not found")
	ErrNotAllowed        = errors.New("method not allowed")
	ErrServerStatusNotOK = errors.New("non-ok server status")
)
