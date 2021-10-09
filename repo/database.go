package repo

import (
	"encoding/csv"
	"os"

	"github.com/valp0/academy-go-q32021/common"
)

type Repo struct{}

// Returns the the element matching the provided ID.
func (r Repo) GetElements(id, path string) ([]common.Element, error) {
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

	if id == "" {
		elements, err := getElements(content)
		if err != nil {
			return []common.Element{}, err
		}

		return elements, nil
	}

	elements, err := filterElement(id, content)
	if err != nil {
		return []common.Element{}, err
	}

	return elements, nil
}
