package handlers

import (
	"fmt"
	"net/http"

	"github.com/valp0/academy-go-q32021/pkg/utils"
)

// The entry point handler.
func (hh homeHandler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jRes, _ := utils.JsonResponse("At this stage, the only available endpoint is /read.")
	prettified := utils.Prettify(jRes)

	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, prettified)
	if err != nil {
		fmt.Println(err)
		return
	}
}
