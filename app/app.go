package app

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/valp0/academy-go-q32021/handlers"
	"github.com/valp0/academy-go-q32021/repo"
	"github.com/valp0/academy-go-q32021/services"
)

const (
	port = ":8080"
	path = "./files/pokemons.csv"
)

func RunServer() error {
	os.Setenv("PATH", path)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go logExit(c)

	r := repo.Repo{}

	rSvc := services.NewReadSvc(r)
	fSvc := services.NewFetchSvc(r)

	hh := handlers.NewHomeHandler()
	rh := handlers.NewReadHandler(rSvc)
	fh := handlers.NewFetchHandler(fSvc)

	http.HandleFunc("/", hh.Home)
	http.HandleFunc("/read", rh.Read)
	http.HandleFunc("/fetch", fh.ApiFetch)

	log.Println("Listening on port", port[1:])
	if err := http.ListenAndServe(port, nil); err != nil {
		return err
	}

	return nil
}

func logExit(c chan os.Signal) {
	for range c {
		log.Fatal("Process terminated")
	}
}
