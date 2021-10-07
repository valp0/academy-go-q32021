package services

import (
	"errors"
	"strconv"

	"github.com/valp0/academy-go-q32021/common"
	"github.com/valp0/academy-go-q32021/repo"
)

// Read ID query param from fetch handler
// Fetch and write pokemon with matching ID to csv (getById.go repo)
// Write random pokemon
// Return whole pokemon list (displayAll.go repo)

// Service that reads the id param and adds pokemon with said ID to csv.
func FetchSvc(params map[string][]string) (string, error) {
	id, ok := params["id"]
	if !ok || len(id[0]) < 1 {
		err := pokeApi(nil)
		if err != nil {
			return "", err
		}
	} else {
		err := checkId(id[0])
		if err != nil {
			return "", err
		}
	}

	return repo.DisplayAll()
}

func checkId(id string) error {
	pokeId, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("Invalid ID, ID must be an integer.")
	}

	if 1 <= pokeId && pokeId <= 898 {
		pokeApi(&pokeId)
		return nil
	} else {
		return errors.New("Invalid ID, ID must be between 1 and 898.")
	}
}

func pokeApi(id *int) error {
	var idInt int
	if id == nil {
		idInt = common.RandInt(899)
	} else {
		idInt = *id
	}

	url := "https://pokeapi.co/api/v2/pokemon-form/" + strconv.Itoa(idInt)
	return repo.CallPokeApi(url)
}
