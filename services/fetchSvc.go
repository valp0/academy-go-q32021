package services

import (
	"errors"
	"strconv"

	"github.com/valp0/academy-go-q32021/common"
)

// Read ID query param from fetch handler
// Fetch and write pokemon with matching ID to csv (getById.go repo)
// Write random pokemon
// Return whole pokemon list (displayAll.go repo)

type apiCaller interface {
	CallPokeApi(url, path string) ([]common.Element, error)
}

type fetchSvc struct {
	repo repository
}

func NewFetchSvc(repo repository) fetchSvc {
	return fetchSvc{repo}
}

// Service that reads the id param and adds pokemon with said ID to csv.
func (fs fetchSvc) Fetch(params map[string][]string, path string) ([]common.Element, error) {
	var sentId *string
	id, ok := params["id"]
	if ok {
		sentId = &id[0]
	} else {
		sentId = nil
	}

	pokeId, err := checkId(sentId, fs)
	if err != nil {
		return []common.Element{}, err
	}

	res, err := pokeApi(fs, pokeId, path)
	if err != nil {
		return []common.Element{}, err
	}

	return res, nil
}

func checkId(id *string, fs fetchSvc) (*int, error) {
	if id == nil {
		return nil, nil
	}

	pokeId, err := strconv.Atoi(*id)
	if err != nil {
		return &pokeId, errors.New("invalid id, id must be an integer")
	}

	if 1 <= pokeId && pokeId <= 898 {
		return &pokeId, nil
	} else {
		return &pokeId, errors.New("invalid id, id must be between 1 and 898")
	}
}

func pokeApi(fs fetchSvc, id *int, path string) ([]common.Element, error) {
	var idInt int
	if id == nil {
		idInt = common.RandInt(899)
	} else {
		idInt = *id
	}

	url := "https://pokeapi.co/api/v2/pokemon-form/" + strconv.Itoa(idInt)
	return fs.repo.CallPokeApi(url, path)
}
