package repo

import (
	"encoding/csv"
	"os"

	"github.com/valp0/academy-go-q32021/common"
)

type localRepo struct{}

// Returns a local repository to handle local database.
func NewLocalRepo() localRepo {
	return localRepo{}
}

// Returns the the element matching the provided ID.
func (r localRepo) GetElements(id, path string) ([]common.Element, error) {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return []common.Element{}, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	content, err := csvReader.ReadAll()
	if err != nil {
		return []common.Element{}, err
	}

	elements, err := getElements(content)
	if err != nil {
		return []common.Element{}, err
	}

	if id == "" {
		sortElements(elements)
		return elements, nil
	}

	element, err := filterElement(id, elements)
	if err != nil {
		return []common.Element{}, err
	}

	return element, nil
}
