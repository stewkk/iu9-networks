package dota2ru

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/stewkk/iu9-networks/lab2/internal/accesslog"
	"github.com/stewkk/iu9-networks/lab2/internal/errors"
)

func RegisterHandlers(mux *httprouter.Router, service Service, logger *log.Logger) {
	res := resource{service, logger}

	mux.GET("/", accesslog.Log(errors.Handle(res.handleGet, logger), logger))
	mux.GET("/:page", accesslog.Log(errors.Handle(res.handleGet, logger), logger))
}

type resource struct {
	service Service
	logger  *log.Logger
}

func (res *resource) handleGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	pageParam := p.ByName("page")
	if pageParam == "" {
		pageParam = "1"
	}
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		return fmt.Errorf("%w: %v", errors.ErrPathParameter, err)
	}

	headings, err := res.service.ParseHeadings(page)
	if err != nil {
		return err
	}

	tmpl, err := template.New("Headings").Parse(`<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Doka</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-Zenh87qX5JnK2Jl0vWa8Ck2rdkQ2Bzep5IDxbcnCeuOxjzrPF/et3URy9Bv1WTRi" crossorigin="anonymous">
  </head>
  <body>
	<div class="container text-center">
        <h1> Железо, новости и обсуждения </h1>
	</div>
	<hr>
	<div class="container-fluid text-start">
		{{range . }}
		<div class="row">
			<div class="col-2">
			</div>
			<div class="col-6">
				<a href="{{.Link}}">{{.Title}}</a>
			</div>
		</div>
		{{end}}
	</div>
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.2/dist/js/bootstrap.min.js" integrity="sha384-IDwe1+LCz02ROU9k972gdyvl+AESN10+x7tBKgc9I5HFtuNz0wWnPclzo6p9vxnk" crossorigin="anonymous"></script>
  </body>
</html>
`)

	tmpl.ExecuteTemplate(w, "Headings", headings)
	return nil
}
