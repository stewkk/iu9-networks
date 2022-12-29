package main

import (
	"fmt"
	"io"
	"log"

	"github.com/stewkk/iu9-networks/lab6/internal/client"
)

func main() {
	client, err := client.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	handlers := map[string]func(string)error{
		"store": client.Store,
		"get": client.Get,
		"mkdir": client.Mkdir,
		"remove": client.Remove,
		"ls": client.Ls,
	}

	for {
		fmt.Print("> ")
		var cmd, path string
		_, err := fmt.Scan(&cmd, &path)
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal(err)
		}

		handler := handlers[cmd]
		if handlers == nil {
			log.Fatal("")
		}

		err = handler(path)
		if err != nil {
			log.Println(err)
		}
	}
}
