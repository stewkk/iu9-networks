package main

import (
	"io"
	"os"

	"github.com/stewkk/iu9-networks/lab4/internal/pty"
	"golang.org/x/term"
)

func main() {
	rw, err := pty.NewPty("/bin/bash")
	if err != nil {
		panic(err)
	}
	defer rw.Close()
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	go func() {
		io.Copy(rw, os.Stdout)
	}()
	io.Copy(os.Stdin, rw)
}
