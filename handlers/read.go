package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/valp0/academy-go-q32021/common"
)

// Use readSvc.go service to handle received params
// Take ID as query param to display one pokemon from file
// Display all pokemons in csv if no ID is given

type reader interface {
	Read(params map[string][]string, path string) ([]common.Element, error)
}

type readHandler struct {
	service reader
}

// Receives an instance of a type that satisfies the reader interface
// and returns a readHandler type containing it.
func NewReadHandler(service reader) readHandler {
	return readHandler{service}
}

const (
	invalidId   = "please use integers for ID field"
	invalidPath = "encontrar la ruta especificada."
	noIdFound   = "wasFound"
)

// The /read endpoint handler.
func (rh readHandler) Query(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		common.MethodNotAllowedError(w, r.Method)
		return
	}

	path := os.Getenv("PATH")

	qParams := r.URL.Query()
	checkMsg := strings.HasSuffix
	res, err := rh.service.Read(qParams, path)
	if err != nil {
		msg := err.Error()
		switch {
		case checkMsg(msg, invalidId), checkMsg(msg, invalidPath), checkMsg(msg, noIdFound):
			common.BadReqError(w, err)
			return
		default:
			common.InternalError(w, err)
			return
		}
	}

	b, err := fmt.Fprintln(w, common.PrettyJsonRes(res))
	if err != nil || b < 1 {
		common.InternalError(w, err)
	}
}
