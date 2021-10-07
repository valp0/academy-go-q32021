package handlers

import (
	"fmt"
	"net/http"

	"github.com/valp0/academy-go-q32021/common"
	"github.com/valp0/academy-go-q32021/services"
)

// Use readSvc.go service to handle received params
// Take ID as query param to display one pokemon from file
// Display all pokemons in csv if no ID is given

type readHandler struct{}

type iReadHandler interface {
	Read(w http.ResponseWriter, r *http.Request)
}

func NewReadHandler() iReadHandler {
	return readHandler{}
}

// The /read endpoint handler.
func (readHandler) Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	qParams := r.URL.Query()
	resp, err := services.ReadSvc(qParams)
	if err != nil {
		common.InternalError(w, err)
	}

	_, err = fmt.Fprintln(w, resp)
	if err != nil {
		common.InternalError(w, err)
	}
}
