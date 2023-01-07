package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/stewkk/iu9-networks/lab8/pkg/reports"
)

const (
	build = "/tmp/lab8"
)

func main() {
	os.Mkdir("/tmp/lab8", 0777)
	root := os.Getenv("LAB8_ROOT")

	server, err := NewServer(root)
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
    router.GET("/*filepath", server.HandleGet)
    router.POST("/", server.HandlePost)

	log.Fatal(http.ListenAndServe(":8080", router))
}

type Server struct {
	root string
	fsHandler http.Handler
}

func NewServer(root string) (*Server, error) {
	return &Server{
		root:      root,
		fsHandler: http.FileServer(http.Dir(root)),
	}, nil
}

func (s *Server) HandleGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	filepath := p.ByName("filepath")
	if path.Ext(filepath) == ".md" {
		id := uuid.New()
		pdf := fmt.Sprintf("%s/%s.pdf", build, id)
		output, err := reports.CompileMarkdown(s.root+filepath, pdf)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(output)
			return
		}
		http.ServeFile(w, r, pdf)
		return
	}

	s.fsHandler.ServeHTTP(w, r)
}

func (s *Server) HandlePost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprint(w, "TODO POST\n")
}
