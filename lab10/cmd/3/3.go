package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/stewkk/iu9-networks/lab3/pkg/kommersant"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := httprouter.New()
	router.GET("/", WsHandler)

	s := &http.Server{
		Handler: router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	l, err := net.Listen("tcp", ":8083")
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
		news, err := kommersant.List()
		if err != nil {
			return err
		}

		err = wsjson.Write(context.TODO(), c, struct {
			Entries []kommersant.NewsEntry `json:"entries"`
		}{
			Entries: news,
		})
		if err != nil {
			return err
		}

		time.Sleep(time.Second)
	}
}
