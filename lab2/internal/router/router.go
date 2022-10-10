package router

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/stewkk/iu9-networks/lab2/internal/accesslog"
	"github.com/stewkk/iu9-networks/lab2/internal/errors"
)

// New returns preconfigured httprouter.Router.
func New(logger *log.Logger) *httprouter.Router {
	return &httprouter.Router{
		HandleMethodNotAllowed: true,
		NotFound: toStdHandler(accesslog.Log(errors.Handle(returnNotFound, logger), logger)),
		MethodNotAllowed: toStdHandler(accesslog.Log(errors.Handle(returnNotAllowed, logger), logger)),
	}
}

func returnNotFound(http.ResponseWriter, *http.Request, httprouter.Params) error {
	return errors.ErrNotFound
}

func returnNotAllowed(http.ResponseWriter, *http.Request, httprouter.Params) error {
	return errors.ErrNotAllowed
}

func toStdHandler(handler httprouter.Handle) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, nil)
	})
}
