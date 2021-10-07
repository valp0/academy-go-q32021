package common

import (
	"fmt"
	"log"
	"net/http"
)

// Set a received error to a received reponse writer
// Write error to response as JSON message
// Log error to console

func InternalError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	displayErr(w, err)
}

func ExternalError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusServiceUnavailable)
	displayErr(w, err)
}

func displayErr(w http.ResponseWriter, err error) {
	jErr := JsonResponse(err.Error())
	if _, e := fmt.Fprintln(w, PrettifyJson(jErr)); e != nil {
		log.Println(e)
	}

	log.Println(err)
}
