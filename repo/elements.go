package repo

import (
	"errors"
	"fmt"
	"log"
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
			str := fmt.Sprintf("Parsing error in line %d: “%s” is not a valid ID, please use integers for ID field.",
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
		err := errors.New("“" + id + "” is not a valid ID, only integer type is allowed as an ID.")
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

	err = errors.New("No element with ID “" + id + "” was found.")
	return []common.Element{}, err
}

func sortElements(elements []common.Element) {
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].Key < elements[j].Key
	})
}
