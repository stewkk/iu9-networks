package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/pbar1/pkill-go"
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

	l, err := net.Listen("tcp", ":8084")
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
	for {
		pids, err := pkill.Pgrep("^ping$")
		if err != nil {
			return err
		}

		var output string
		switch len(pids) {
		case 0:
			output = "no pings to BAUMAN"
		case 1:
			output = "someone is pinging bmstu.ru"
		default:
			output = fmt.Sprintf("%v pings to BAUMAN!", len(pids))
		}
		err = wsjson.Write(context.TODO(), c, struct {
			Output string `json:"output"`
		}{
			Output: output,
		})
		if err != nil {
			return err
		}

		time.Sleep(time.Second)
	}
}
