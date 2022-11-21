package main

import (
	"context"
	"fmt"
	"time"

	"github.com/stewkk/iu9-networks/lab5/internal/api"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://localhost:5332", nil)
	if err != nil {
		panic(err)
	}
	defer c.Close(websocket.StatusInternalError, "internal error")

	var integral api.Input
	fmt.Scan(&integral.A, &integral.B, &integral.C, &integral.Start, &integral.End)

	err = wsjson.Write(ctx, c, integral)
	if err != nil {
		panic(err)
	}

	var result api.Result
	err = wsjson.Read(ctx, c, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.Sum)

	c.Close(websocket.StatusNormalClosure, "")
}
