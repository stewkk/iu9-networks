package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"os/exec"
	"time"

	"github.com/julienschmidt/httprouter"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func main() {
	log.Fatal(runServer())
}

func runServer() error {
	server := Server{
	}
	router := httprouter.New()
	router.GET("/", server.IndexHandler)
	router.GET("/subscribe", server.SubscribeHandler)

	s := &http.Server{
		Handler: router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		return err
	}
	log.Printf("listening on http://%v", l.Addr())

	return s.Serve(l)
}

type Server struct {}

func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "html/async.html")
}

func (s *Server) SubscribeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "")

	err = s.subscribe(c)
	if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
		websocket.CloseStatus(err) == websocket.StatusGoingAway {
		return
	}
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *Server) subscribe(c *websocket.Conn) error {
	cmd := exec.Command("sh", "-s")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	out := io.MultiReader(stdout, stderr)
	err = cmd.Start()
	if err != nil {
		return err
	}

	errc := make(chan error, 1)
	go func() {
		scanner := bufio.NewScanner(out)
		for scanner.Scan() {
			ctx := context.TODO()
			line := scanner.Text()
			log.Println("sending line of output: ", line)
			err := wsjson.Write(ctx, c, Output{
				Line: line+"\n",
			})
			if err != nil {
				errc <- err
			}
		}
	}()

	go func() {
		var input Input
		for {
			ctx := context.TODO()

			err := wsjson.Read(ctx, c, &input)
			if err != nil {
				errc <- err
			}
			log.Println("received input:", input.Cmd)
			stdin.Write([]byte(input.Cmd+"\n"))
		}
	}()

	return <-errc
}

type Input struct {
	Cmd string `json:"cmd"`
}

type Output struct {
	Line string `json:"line"`
}
