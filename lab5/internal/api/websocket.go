package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/stewkk/iu9-networks/lab5/internal/integral"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func RunServer() {
	server := http.Server{
		Addr: "0.0.0.0:5332",
		Handler: http.HandlerFunc(handler),
	}

	log.Println("Listenig...")
	server.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "internal error")

	for {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		var v integral.Integral
		err = wsjson.Read(ctx, c, &v)
		if err != nil {
			fmt.Println(err)
			return
		}

		sum := v.Calc()
		log.Printf("received: %v\tsum: %v", v, sum)

		wsjson.Write(ctx, c, Result{
			Sum: sum,
		})
	}
}
