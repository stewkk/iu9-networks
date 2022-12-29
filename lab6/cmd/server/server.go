package main

import (
	"log"
	"os"

	"goftp.io/server/v2"
	"goftp.io/server/v2/driver/file"
)

func main() {
	os.Mkdir("tmp/", 0777)
	driver, err := file.NewDriver("tmp/")
	auth := &server.SimpleAuth{
		Name:     os.Getenv("LOGIN"),
		Password: os.Getenv("PASSWD"),
	}
	perm := &server.SimplePerm{}
	opts := &server.Options{
		Driver: driver,
		Auth: auth,
		Port: 2000,
		Perm: perm,
		Hostname: "0.0.0.0",
	}
	server, err  := server.NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.ListenAndServe())
}
