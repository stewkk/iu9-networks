package accesslog

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Log middleware logs handled request.
func Log(handler httprouter.Handle, logger *log.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		logger.Println(r.Method, r.URL, r.Proto)
		handler(w, r, p)
	}
}
