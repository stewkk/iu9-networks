package main

import (
	"log"
	"net/http"
	"os"

	"github.com/stewkk/iu9-networks/lab2/internal/dota2ru"
	"github.com/stewkk/iu9-networks/lab2/internal/router"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	mux := router.New(logger)

	dota2ru.RegisterHandlers(mux, dota2ru.NewService(), logger)

	address := "0.0.0.0:80"
	server := http.Server{
		Addr:    address,
		Handler: mux,
	}

	logger.Println("Slavatidika launched on", address)
	logger.Println(server.ListenAndServe())
	os.Exit(1)
}
