package handlers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/valp0/academy-go-q32021/pkg/utils"
)

type element struct {
	Key   int         `json:"id"`
	Value interface{} `json:"name"`
}

// The /read endpoint handler.
func ReadLocalCsv(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fileArr, err := readCsv("./files/pokemons.csv")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jErr, _ := utils.JsonResponse("Couldn't read local csv file. " + err.Error())
		fmt.Fprintln(w, utils.Prettify(jErr))
		return
	}

	pokemonList, err := fillList(fileArr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jErr, _ := utils.JsonResponse(err.Error())
		fmt.Fprintln(w, utils.Prettify(jErr))
		return
	}

	// If we send an id query param, it will be used to fetch pokemons by id.
	id, ok := r.URL.Query()["id"]

	// Will return the entire pokemon list if no id query param is sent.
	if !ok || len(id[0]) < 1 {
		pokeResp, err := utils.JsonResponse(pokemonList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			jErr, _ := utils.JsonResponse("Error generating JSON response. " + err.Error())
			fmt.Fprintln(w, utils.Prettify(jErr))
		}

		w.WriteHeader(http.StatusOK)
		prettified := utils.Prettify(pokeResp)
		fmt.Fprintln(w, prettified)
		return
	}

	pokeSearch, err := queryById(pokemonList, id[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jErr, _ := utils.JsonResponse(err.Error())
		fmt.Fprintln(w, utils.Prettify(jErr))
		return
	}

	pokeResp, err := utils.JsonResponse(pokeSearch)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jErr, _ := utils.JsonResponse("Couldn't read local csv file. " + err.Error())
		fmt.Fprintln(w, utils.Prettify(jErr))
	}

	w.WriteHeader(http.StatusOK)
	prettified := utils.Prettify(pokeResp)
	fmt.Fprintln(w, prettified)
}

// A function that receives a path to a csv file,
// and returns a 2D string slice with its content.
func readCsv(path string) ([][]string, error) {

	// Opening csv file and handling error.
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Unable to open file", path, "\nReason:", err)
		return nil, err
	}
	defer f.Close()

	// Reading from csv file and handling error.
	csvReader := csv.NewReader(f)
	file, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Unable to read file", path, "\nReason:", err)
		return nil, err
	}

	return file, nil
}

// This function will fill a slice of element with the raw 2D received slice.
func fillList(fileArr [][]string) ([]element, error) {
	elements := []element{}

	for i, elem := range fileArr {
		name := elem[1]

		id, err := strconv.Atoi(elem[0])
		if err != nil {
			str := fmt.Sprintf("Parsing error in line %d: “%s” is not a valid ID, please use integers for ID field.",
				i+1, err.Error()[23:23+len(elem[0])])
			fmt.Println(str)
			newErr := errors.New(str)
			return nil, newErr
		}

		elements = append(elements, element{id, name})
	}

	return elements, nil
}

// This function returns the element with specified ID
func queryById(elements []element, id string) ([]element, error) {
	pokeId, err := strconv.Atoi(id)
	if err != nil {
		err := errors.New("“" + id + "” is not a valid ID, only integer type is allowed as an ID.")
		return []element{}, err
	}

	for _, e := range elements {
		if e.Key == pokeId {
			return []element{e}, nil
		}
	}

	err = errors.New("No element with ID “" + id + "” was found.")
	return []element{}, err
}
