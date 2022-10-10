package errors

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Handler type represents handler that can return error. It's intended to be
// wrapped in error.Handle handler
type Handler func(http.ResponseWriter, *http.Request, httprouter.Params) error

// Handle creates middleware for handling errors and panics encountered during request handling
func Handle(handler Handler, logger *log.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer func() {
			if err := recover(); err != nil {
				logger.Println(err)
			}
		}()

		err := handler(w, r, p)
		if err != nil {
			logger.Println(err)
			errorResponse(w, err, logger)
		}
	}
}

func errorResponse(w http.ResponseWriter, err error, logger *log.Logger) {
	switch {
	case errors.Is(err, ErrNotFound):
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not found")
	case errors.Is(err, ErrNotAllowed):
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed")
	case errors.Is(err, ErrPathParameter):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad request")
	default:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal server error")
	}
}
