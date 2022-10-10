package main

import (
	"fmt"
	"os"

	"github.com/stewkk/iu9-networks/lab3/internal/kommersant"
)

func main() {
	news, err := kommersant.ParseFeed()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = kommersant.Store(news)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
