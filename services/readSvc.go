package services

import "github.com/valp0/academy-go-q32021/repo"

// Read ID query param from read handler
// Return whole pokemon list if no ID is given (displayAll.go repo)
// Return only pokemon with matching ID if ID is given (getById.go repo)

// Service that reads the id param and returns a prettified JSON response.
func ReadSvc(params map[string][]string) (string, error) {
	id, ok := params["id"]
	if !ok || len(id[0]) < 1 {
		return repo.DisplayAll()
	}

	return repo.GetById(id[0])
}
