package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/valp0/academy-go-q32021/common"
)

type selector interface {
	Select(params map[string][]string, path string) ([]common.Element, error)
}

type asyncHandler struct {
	service selector
}

// Receives an instance of a type that satisfies the reader interface
// and returns a readHandler type containing it.
func NewAsyncHandler(service selector) asyncHandler {
	return asyncHandler{service}
}

// The /read endpoint handler.
func (ah asyncHandler) Async(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		common.MethodNotAllowedError(w, r.Method)
		return
	}

	path := os.Getenv("PATH")

	qParams := r.URL.Query()
	res, err := ah.service.Select(qParams, path)
	if err != nil {
		common.InternalError(w, err)
		return
	}

	b, err := fmt.Fprintln(w, common.PrettyJsonRes(res))
	if err != nil || b < 1 {
		common.InternalError(w, err)
	}
}
