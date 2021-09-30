package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/valp0/academy-go-q32021/pkg/handlers"
)

const port = ":8080"

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go logExit(c)

	hh := handlers.NewHomeHandler()
	rh := handlers.NewReadHandler()

	http.HandleFunc("/", hh.Home)
	http.HandleFunc("/read", rh.ReadLocalCsv)

	log.Println("Listening on port", port[1:])
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err.Error())
	}
}

func logExit(c chan os.Signal) {
	for range c {
		log.Fatal("Process terminated")
	}
}
