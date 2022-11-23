package main

import (
	"fmt"
	"io"
	"log"

	"github.com/gliderlabs/ssh"
	"github.com/stewkk/iu9-networks/lab4/internal/pty"
)

func main() {
	authHandler := ssh.PasswordAuth(func(ctx ssh.Context, password string) bool {
		return ctx.User() == "test" && password == "12345678"
	})
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
		rw, err := pty.ExecWithPty("/bin/sudo", "--login", "--user", s.User())
		if err != nil {
			panic(err)
		}
		defer rw.Close()
		go func() {
			io.Copy(rw, s)
		}()
		io.Copy(s, rw)
	})

	log.Println("starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil, authHandler))
}
