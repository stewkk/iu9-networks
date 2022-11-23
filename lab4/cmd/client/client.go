package main

import (
	"os"

	"github.com/helloyi/go-sshclient"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

var remoteConfig = &ssh.ClientConfig{
	User:            "test",
	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
}

var localConfig = &ssh.ClientConfig{
	User:            "test",
	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
}

var (
	remote = ""
	local  = "localhost:2222"
)

func main() {
	client, err := sshclient.DialWithPasswd(local, localConfig.User, "12345678")
	if err != nil {
		panic(err)
	}

	// with a terminal config
	config := &sshclient.TerminalConfig{
		Term:   "xterm",
		Height: 40,
		Weight: 80,
		Modes: ssh.TerminalModes{
			ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
			ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
		},
	}

	// Set stdin in raw mode.
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }() // Best effort.

	err = client.Terminal(config).Start()
	if err != nil {
		panic(err)
	}
}
