package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/stewkk/iu9-networks/lab3.2/pkg/inotify"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func main() {
	router := httprouter.New()
	router.GET("/", WsHandler)

	s := &http.Server{
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	l, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening on http://%v", l.Addr())

	log.Fatal(s.Serve(l))
}

func WsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "")

	err = subscribe(c)
	if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
		websocket.CloseStatus(err) == websocket.StatusGoingAway {
		return
	}
	if err != nil {
		log.Println(err)
		return
	}
}

func subscribe(c *websocket.Conn) error {
	watch, err := inotify.NewWatch(".")
	if err != nil {
		return err
	}
	for event := range watch.Events {
		if event.Path != "achtung.txt" {
			continue
		}

		var contents string
		switch event.Type {
		case inotify.CREATE:
			tmp, err := ioutil.ReadFile(event.Path)
			contents = string(tmp)
			if err != nil {
				return err
			}
		case inotify.DELETE:
			contents = "norm"
		}

		err = wsjson.Write(context.TODO(), c, struct {
			Contents string `json:"contents"`
		}{
			Contents: contents,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
