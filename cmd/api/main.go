package main

import (
	"log"
	"net/http"

	"github.com/valp0/academy-go-q32021/pkg/handlers"
)

const port = ":8080"

func main() {
	hh := handlers.NewHomeHandler()
	rh := handlers.NewReadHandler()

	http.HandleFunc("/", hh.Home)
	http.HandleFunc("/read", rh.ReadLocalCsv)

	log.Println("Listening on port", port[1:])
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
