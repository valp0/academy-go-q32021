package handlers

import (
	"fmt"
	"net/http"

	"github.com/valp0/academy-go-q32021/common"
)

// Give a message to inform about available endpoints
// Display message as JSON (use jsonResponse.go common)

type informer interface {
	Inform() string
}

type homeHandler struct {
	service informer
}

// Receives an instance of a type that satisfies the informer interface
// and returns a homeHandler type containing it.
func NewHomeHandler(service informer) homeHandler {
	return homeHandler{service}
}

// The / endpoint handler.
func (hh homeHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		common.MethodNotAllowedError(w, r.Method)
		return
	}

	msg := hh.service.Inform()
	res := common.PrettyJsonRes(msg)

	w.WriteHeader(http.StatusOK)
	b, err := fmt.Fprint(w, res)
	if err != nil || b < 1 {
		common.InternalError(w, err)
	}
}
