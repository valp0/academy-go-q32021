package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/valp0/academy-go-q32021/common"
)

// Use readSvc.go service to handle received params
// Take ID as query param to display one pokemon from file
// Display all pokemons in csv if no ID is given

type query interface {
	Query(params map[string][]string, path string) ([]common.Element, error)
}

type readHandler struct {
	service query
}

func NewReadHandler(service query) readHandler {
	return readHandler{service}
}

// The /read endpoint handler.
func (rh readHandler) Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	path := os.Getenv("PATH")

	qParams := r.URL.Query()
	res, err := rh.service.Query(qParams, path)
	if err != nil {
		if len(err.Error()) > 51 && err.Error()[len(err.Error())-51:] == "is not a valid ID, please use integers for ID field" ||
			len(err.Error()) > 40 && err.Error()[len(err.Error())-40:] == "El sistema no puede encontrar la ruta especificada." ||
			len(err.Error()) > 9 && err.Error()[len(err.Error())-9:] == "was found" {
			common.BadReqError(w, err)
			return
		}

		common.InternalError(w, err)
		return
	}

	_, err = fmt.Fprintln(w, common.PrettyJsonRes(res))
	if err != nil {
		common.InternalError(w, err)
	}
}
