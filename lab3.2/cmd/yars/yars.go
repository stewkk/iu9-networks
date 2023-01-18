package main

import (
	"fmt"

	"github.com/stewkk/iu9-networks/lab3.2/internal/inotify"
)

func main() {
	watch, err := inotify.NewWatch(".")
	if err != nil {
		panic(err)
	}

	for event := range watch.Events {
		fmt.Println(event)
	}
}
