package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/stewkk/iu9-networks/lab9/pkg/terminals"
)

func main() {
	tmpl, err := template.New("index").ParseFiles("html/sync.html")
	if err != nil {
		log.Fatal(err)
	}

	server := Server{
		tmpl: tmpl,
	}

	router := httprouter.New()
	router.GET("/", server.GetHandler)
	router.POST("/", server.PostHandler)

	log.Fatal(http.ListenAndServe(":8081", router))
}

type Server struct {
	tmpl *template.Template
}

func (s *Server) GetHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := s.tmpl.ExecuteTemplate(w, "index", "")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Server) PostHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cmd := r.PostFormValue("cmd")

	output := terminals.Cmd(cmd)

	err := s.tmpl.ExecuteTemplate(w, "index", output)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
