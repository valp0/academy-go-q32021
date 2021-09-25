package handlers

import (
	"fmt"
	"net/http"

	"github.com/valp0/academy-go-q32021/pkg/utils"
)

type homeHandler struct{}

type IHomeHandler interface {
	Home(w http.ResponseWriter, r *http.Request)
}

// Will return homeHandler struct.
func NewHomeHandler() IHomeHandler {
	return homeHandler{}
}

// The entry point handler.
func (hh homeHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jRes := utils.JsonResponse("At this stage, the only available endpoint is /read.")
	prettified := utils.Prettify(jRes)

	w.WriteHeader(http.StatusOK)
	b, err := fmt.Fprintf(w, prettified)
	if err != nil || b < 1 {
		fmt.Println("Couldn't write response bytes to http.ResponseWriter,", err)
	}
}
