package common

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Will set an internal server status (500 code) to header and write
// received error to received response writer and console.
func InternalError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	displayErr(w, err)
}

// Will set a service unavailable status (503 code) to header and write
// received error to received response writer and console.
func ExternalError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusServiceUnavailable)
	displayErr(w, err)
}

// Will set a bad request status (400 code) to header and write
// received error to received response writer and console.
func BadReqError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	displayErr(w, err)
}

// Will set a method not allowed status (405 code) to header and write
// received method name to received response writer and console.
func MethodNotAllowedError(w http.ResponseWriter, method string) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	err := errors.New(fmt.Sprintf("method %s not allowed on this endpoint", method))
	displayErr(w, err)
}

func displayErr(w http.ResponseWriter, err error) {
	jErr := JsonResponse(err.Error())
	if _, e := fmt.Fprintln(w, PrettifyJson(jErr)); e != nil {
		log.Println(e)
	}

	log.Println(err)
}
