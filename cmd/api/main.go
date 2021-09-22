package main

import (
	"log"
	"net/http"

	"github.com/valp0/academy-go-q32021/pkg/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/read", handlers.ReadLocalCsv)

	log.Println("Listening on port", port[1:])
	err := http.ListenAndServe(port, nil)
	if err == nil {
		log.Fatal(err)
	}
}
