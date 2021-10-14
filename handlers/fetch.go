package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/valp0/academy-go-q32021/common"
)

// Call API to fetch new pokemon and write it to file
// Take ID as query param to fetch info from a specific pokemon
// Fetch random pokemon info if no ID is given
// Write said info to csv file

type fetcher interface {
	Fetch(params map[string][]string, path string) ([]common.Element, error)
}

type fetchHandler struct {
	service fetcher
}

// Receives an instance of a type that satisfies the fetch interface
// and returns a fetchHandler type containing it.
func NewFetchHandler(service fetcher) fetchHandler {
	return fetchHandler{service}
}

const (
	notInt     = "must be an integer"
	notInRange = "between 1 and 898"
	existingId = "already stored in csv"
)

// The /fetch endpoint handler.
func (fh fetchHandler) ApiFetch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		common.MethodNotAllowedError(w, r.Method)
		return
	}

	path := os.Getenv("PATH")

	qParams := r.URL.Query()
	checkMsg := strings.HasSuffix
	res, err := fh.service.Fetch(qParams, path)
	if err != nil {
		msg := err.Error()
		switch {
		case checkMsg(msg, notInRange), checkMsg(msg, notInt):
			common.BadReqError(w, err)
			return
		case checkMsg(msg, existingId):
			common.InternalError(w, err)
			return
		default:
			common.ExternalError(w, err)
			return
		}
	}

	b, err := fmt.Fprintln(w, common.PrettyJsonRes(res))
	if err != nil || b < 1 {
		common.InternalError(w, err)
	}
}
