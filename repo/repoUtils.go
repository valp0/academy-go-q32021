package repo

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/valp0/academy-go-q32021/common"
)

func getElements(content [][]string) ([]common.Element, error) {
	elements := []common.Element{}
	for i, elem := range content {
		name := elem[1]

		id, err := strconv.Atoi(elem[0])
		if err != nil {
			str := fmt.Sprintf("parsing error in line %d: \"%s\" is not a valid id, please use integers for id field",
				i+1, err.Error()[23:23+len(elem[0])])
			log.Println(str)
			newErr := errors.New(str)
			return []common.Element{}, newErr
		}

		elements = append(elements, common.Element{Key: id, Value: name})
	}

	sortElements(elements)
	return elements, nil
}

func filterElement(id string, content [][]string) ([]common.Element, error) {
	pokeId, err := strconv.Atoi(id)
	if err != nil {
		err := errors.New("only integer type is allowed as an id")
		return []common.Element{}, err
	}

	elements, err := getElements(content)
	if err != nil {
		return []common.Element{}, err
	}

	for _, e := range elements {
		if e.Key == pokeId {
			return []common.Element{e}, nil
		}
	}

	err = errors.New("no element with id " + id + " was found")
	return []common.Element{}, err
}

func sortElements(elements []common.Element) {
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].Key < elements[j].Key
	})
}

func alreadyStored(id int, path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	content, err := csvReader.ReadAll()
	if err != nil {
		return false, err
	}

	elements, err := getElements(content)
	if err != nil {
		return false, err
	}

	for _, element := range elements {
		if element.Key == id {
			return true, nil
		}
	}

	return false, nil
}
