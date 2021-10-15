package services

import "github.com/valp0/academy-go-q32021/common"

// Read ID query param from read handler
// Return whole pokemon list if no ID is given (displayAll.go repo)
// Return only pokemon with matching ID if ID is given (getById.go repo)

type filter interface {
	Filter(t, items, ipw, path string) ([]common.Element, error)
}

type asyncSvc struct {
	repo filter
}

// Receives an instance of a type that satisfies the getter interface
// and returns a readSvc type containing it.
func NewAsyncSvc(repo filter) asyncSvc {
	return asyncSvc{repo}
}

// Receives a url.Values, reads the type, items and items_per_worker, and returns a prettified JSON response.
func (rs asyncSvc) Select(params map[string][]string, path string) ([]common.Element, error) {
	t, ok := params["type"]
	if !ok {
		t = []string{""}
	}

	items, ok := params["items"]
	if !ok {
		items = []string{""}
	}

	ipw, ok := params["items_per_worker"]
	if !ok {
		ipw = []string{""}
	}

	res, err := rs.repo.Filter(t[0], items[0], ipw[0], path)
	if err != nil {
		return []common.Element{}, err
	}

	return res, nil
}
