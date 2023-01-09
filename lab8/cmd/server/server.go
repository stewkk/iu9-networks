package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/stewkk/iu9-networks/lab8/pkg/reports"
)

func main() {
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
	generator reports.ReportGenerator
}

func NewServer(root string) (*Server, error) {
	gen, err := reports.NewReportGenerator()
	if err != nil {
		return nil, err
	}
	return &Server{
		root:      root,
		fsHandler: http.FileServer(http.Dir(root)),
		generator: gen,
	}, nil
}

func (s *Server) HandleGet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	filepath := p.ByName("filepath")
	if path.Ext(filepath) == ".md" {
		id := uuid.New()
		pdf := fmt.Sprintf("%s/%s.pdf", s.root, id)
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
	var report reports.Fields
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	report.Year = strconv.Itoa(time.Now().Year())


	id := uuid.New()
	title := fmt.Sprintf("%s/%s.title", s.root, id)
	output, err := s.generator.GenerateTitle(&report, title)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(output))
		return
	}
	body := fmt.Sprintf("%s/%s.body", s.root, id)
	bodyFile, err := os.Create(body+".md")
	_, err = fmt.Fprint(bodyFile, report.Body)
	output, err = reports.CompileMarkdown(body+".md", body+".pdf")
	res := fmt.Sprintf("%s/%s.pdf", s.root, id)
	output, err = reports.MergePdfs(title+".pdf", body+".pdf", res)

	err = json.NewEncoder(w).Encode(struct {
		Url string `json:"url"`
	}{
		Url: fmt.Sprintf("%s.pdf", id),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
