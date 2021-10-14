package repo

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/valp0/academy-go-q32021/common"
)

type apiRepo struct{}

// Returns an API repository to handle api calls.
func NewApiRepo() apiRepo {
	return apiRepo{}
}

// Receives a url to call an API and a path to the file to store received data.
func (r apiRepo) CallPokeApi(url, path string) ([]common.Element, error) {
	// Calling GET HTTP method, which returns a response and an error.
	resp, err := CallApi(url)

	var pokeId int
	var pokeName string

	if err != nil {
		return []common.Element{}, err
	} else {
		// Closing response body if there is one.
		defer resp.Body.Close()

		// Reading response if HTTP request was successful.
		if resp.StatusCode == 200 {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			respMap := map[string]interface{}{}
			err = json.Unmarshal(bodyBytes, &respMap)
			if err != nil {
				return []common.Element{}, err
			}

			pokeId, err = strconv.Atoi(fmt.Sprintf("%v", respMap["id"]))
			if err != nil {
				return []common.Element{}, err
			}

			pokeName = fmt.Sprintf("%v", respMap["name"])

			stored, err := addPokemon(pokeId, pokeName, path)
			if err != nil {
				return []common.Element{}, err
			}

			if !stored {
				errStr := "pokémon with id " + strconv.Itoa(pokeId) + " was already stored in csv"
				return []common.Element{}, errors.New(errStr)
			} else {
				log.Printf("Added pokémon %s with id %d to csv", pokeName, pokeId)
			}
		} else {
			return []common.Element{}, fmt.Errorf("%s -> status code: %d ", url, resp.StatusCode)
		}
	}

	return []common.Element{{Key: pokeId, Value: pokeName}}, nil
}

func addPokemon(id int, name, path string) (bool, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return false, err
	}

	defer f.Close()

	alreadyStored, err := alreadyStored(id, path)
	if err != nil {
		return false, err
	}

	if alreadyStored {
		return false, nil
	}

	csvWriter := csv.NewWriter(f)
	pokemon := []string{strconv.Itoa(id), name}

	err = csvWriter.Write(pokemon)
	if err != nil {
		return false, err
	}

	csvWriter.Flush()
	return true, nil
}

// Makes an HTTP get request to the given URL and returns its response.
var CallApi = func(url string) (*http.Response, error) {
	return http.Get(url)
}
