package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

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

const (
	negativeItems = "cannot be a negative number"
	wrongParity   = "can only be \"odd\" or \"even\""
	manyFields    = "can only have two fields"
)

// The /read endpoint handler.
func (ah asyncHandler) Async(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		common.MethodNotAllowedError(w, r.Method)
		return
	}

	path := os.Getenv("PATH")

	qParams := r.URL.Query()
	checkMsg := strings.HasSuffix
	res, err := ah.service.Select(qParams, path)
	if err != nil {
		msg := err.Error()
		switch {
		case checkMsg(msg, negativeItems), checkMsg(msg, wrongParity):
			common.BadReqError(w, err)
			return
		case checkMsg(msg, manyFields):
			common.InternalError(w, err)
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
