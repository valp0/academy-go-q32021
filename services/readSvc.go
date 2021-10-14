package services

import "github.com/valp0/academy-go-q32021/common"

// Read ID query param from read handler
// Return whole pokemon list if no ID is given (displayAll.go repo)
// Return only pokemon with matching ID if ID is given (getById.go repo)

type getter interface {
	GetElements(id, path string) ([]common.Element, error)
}

type readSvc struct {
	repo getter
}

// Receives an instance of a type that satisfies the getter interface
// and returns a readSvc type containing it.
func NewReadSvc(repo getter) readSvc {
	return readSvc{repo}
}

// Receives a url.Values, reads the id param and returns a prettified JSON response.
func (rs readSvc) Read(params map[string][]string, path string) ([]common.Element, error) {
	id, ok := params["id"]
	if !ok {
		id = []string{""}
	}

	res, err := rs.repo.GetElements(id[0], path)
	if err != nil {
		return []common.Element{}, err
	}

	return res, nil
}
