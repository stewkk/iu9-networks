package main

import (
	"io"
	"os"

	"golang.org/x/crypto/ssh"
)

var remoteConfig = &ssh.ClientConfig{
	User: "test",
	Auth: []ssh.AuthMethod{
		ssh.Password(""),
	},
	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
}

var localConfig = &ssh.ClientConfig{
	User: "",
	Auth: []ssh.AuthMethod{
		ssh.Password(""),
	},
	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
}

var (
	remote = ""
	local = "localhost:2222"
)

func main() {
	client, err := ssh.Dial("tcp", local, localConfig)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	stdin, err := session.StdinPipe()
	if err != nil {
		panic(err)
	}

	err = session.Shell()
	if err != nil {
		panic(err)
	}

	io.Copy(stdin, os.Stdin)
}
