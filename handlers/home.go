package handlers

import (
	"fmt"
	"net/http"

	"github.com/valp0/academy-go-q32021/common"
)

// Give a message to inform about available endpoints
// Display message as JSON (use jsonResponse.go common)

type homeHandler struct{}

type iHomeHandler interface {
	Home(w http.ResponseWriter, r *http.Request)
}

func NewHomeHandler() iHomeHandler {
	return homeHandler{}
}

// The / endpoint handler.
func (homeHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jRes := common.JsonResponse("At this stage, available endpoints are /read and /fetch.")
	prettified := common.PrettifyJson(jRes)

	w.WriteHeader(http.StatusOK)
	b, err := fmt.Fprintf(w, prettified)
	if err != nil || b < 1 {
		common.InternalError(w, err)
	}
}
