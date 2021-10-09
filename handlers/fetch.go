package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/valp0/academy-go-q32021/common"
)

// Call API to fetch new pokemon and write it to file
// Take ID as query param to fetch info from a specific pokemon
// Fetch random pokemon info if no ID is given
// Write said info to csv file

type fetch interface {
	Fetch(params map[string][]string, path string) ([]common.Element, error)
}

type fetchHandler struct {
	service fetch
}

func NewFetchHandler(service fetch) fetchHandler {
	return fetchHandler{service}
}

// The /fetch endpoint handler.
func (fh fetchHandler) ApiFetch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	path := os.Getenv("PATH")

	qParams := r.URL.Query()
	res, err := fh.service.Fetch(qParams, path)
	if err != nil {
		if err.Error() == "invalid id, id must be an integer" || err.Error() == "invalid id, id must be between 1 and 898" {
			common.BadReqError(w, err)
			return
		}

		if err.Error()[len(err.Error())-25:] == "was already stored in csv" {
			common.InternalError(w, err)
			return
		}

		common.ExternalError(w, err)
		return
	}

	_, err = fmt.Fprintln(w, common.PrettyJsonRes(res))
	if err != nil {
		common.InternalError(w, err)
	}
}
