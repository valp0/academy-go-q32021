package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func CallPokeApi(url string) error {
	// Calling GET HTTP method, which returns a response and an error.
	resp, err := http.Get(url)

	if err != nil {
		return err
	} else {
		// Closing response body if there is one.
		defer resp.Body.Close()

		// Reading response if HTTP request was successful.
		if resp.StatusCode == 200 {
			bodyBytes, _ := ioutil.ReadAll(resp.Body)
			respMap := map[string]interface{}{}
			err = json.Unmarshal(bodyBytes, &respMap)
			if err != nil {
				return err
			}

			pokeId, err := strconv.Atoi(fmt.Sprintf("%v", respMap["id"]))
			if err != nil {
				return err
			}

			pokeName := fmt.Sprintf("%v", respMap["name"])

			err = addPokemon(pokeId, pokeName)
			if err != nil {
				return err
			}

			log.Printf("Added pokemon %s with ID %d to csv", pokeName, pokeId)
		} else {
			return errors.New(fmt.Sprintf("%s -> Status Code: %d  \n", url, resp.StatusCode))
		}
	}

	return nil
}
