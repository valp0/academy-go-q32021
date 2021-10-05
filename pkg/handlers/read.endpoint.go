package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/valp0/academy-go-q32021/pkg/entities"
	"github.com/valp0/academy-go-q32021/pkg/repo"
	"github.com/valp0/academy-go-q32021/pkg/utils"
)

type CSV entities.CSV
type readHandler struct{}

type IReadHandler interface {
	ReadLocalCsv(w http.ResponseWriter, r *http.Request)
}

func NewReadHandler() IReadHandler {
	return readHandler{}
}

const path = "./files/pokemons.csv"

// The /read endpoint handler.
func (readHandler) ReadLocalCsv(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	f, err := repo.NewCsvService(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jErr := utils.JsonResponse("Couldn't read local csv file. " + err.Error())
		b, err := fmt.Fprintln(w, utils.Prettify(jErr))
		if err != nil || b < 1 {
			log.Println("Couldn't write response bytes to http.ResponseWriter,", err)
		}
		return
	}

	// If we send an id query param, it will be used to fetch pokemons by id.
	id, ok := r.URL.Query()["id"]

	// Will return the entire pokemon list if no id query param is sent.
	if !ok || len(id[0]) < 1 {
		pokeResp := utils.JsonResponse(f.Elements)
		w.WriteHeader(http.StatusOK)
		prettified := utils.Prettify(pokeResp)
		b, err := fmt.Fprintln(w, prettified)
		if err != nil || b < 1 {
			log.Println("Couldn't write response bytes to http.ResponseWriter,", err)
		}
		return
	}

	match, err := f.QueryById(id[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jErr := utils.JsonResponse(err.Error())
		b, err := fmt.Fprintln(w, utils.Prettify(jErr))
		if err != nil || b < 1 {
			log.Println("Couldn't write response bytes to http.ResponseWriter,", err)
		}
		return
	}

	resp := utils.JsonResponse(match)
	w.WriteHeader(http.StatusOK)
	prettified := utils.Prettify(resp)
	b, err := fmt.Fprintln(w, prettified)
	if err != nil || b < 1 {
		log.Println("Couldn't write response bytes to http.ResponseWriter,", err)
	}
}
