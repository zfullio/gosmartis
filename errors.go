package smartis

import "errors"

type APIError struct {
	Msg string `json:"error"`
}

func (e *APIError) Error() string {
	return e.Msg
}

var (
	errInternalError = errors.New("internal error")
	errUnauthorized  = errors.New("unauthorized")
)
