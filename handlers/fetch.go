package handlers

import (
	"fmt"
	"net/http"

	"github.com/valp0/academy-go-q32021/common"
	"github.com/valp0/academy-go-q32021/services"
)

// Call API to fetch new pokemon and write it to file
// Take ID as query param to fetch info from a specific pokemon
// Fetch random pokemon info if no ID is given
// Write said info to csv file

type fetchHandler struct{}

type iFetchHandler interface {
	Fetch(w http.ResponseWriter, r *http.Request)
}

func NewFetchHandler() iFetchHandler {
	return fetchHandler{}
}

// The /fetch endpoint handler.
func (fetchHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	qParams := r.URL.Query()
	resp, err := services.FetchSvc(qParams)
	if err != nil {
		common.ExternalError(w, err)
	}

	_, err = fmt.Fprintln(w, resp)
	if err != nil {
		common.InternalError(w, err)
	}
}
