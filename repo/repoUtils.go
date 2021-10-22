package repo

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
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

	return elements, nil
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

func filterElement(id string, elements []common.Element) ([]common.Element, error) {
	pokeId, err := strconv.Atoi(id)
	if err != nil {
		err := errors.New("only integer type is allowed as an id")
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

func checkType(r string) (string, error) {
	if r == "" {
		return "all", nil
	}

	if r == "odd" || r == "even" {
		return r, nil
	} else {
		err := errors.New("type can only be \"odd\" or \"even\"")
		return "", err
	}
}

func checkItems(r string) (int, error) {
	if r == "" {
		return 1000, nil
	}

	if num, err := strconv.Atoi(r); err == nil {
		if num < 0 {
			err := errors.New("items cannot be a negative number")
			return 0, err
		}
		return num, nil
	} else {
		err = errors.New("items can only be integer type")
		return 0, err
	}
}

func checkIpw(r string) (int, error) {
	if r == "" {
		return 500, nil
	}

	if num, err := strconv.Atoi(r); err == nil {
		if num < 0 {
			err := errors.New("items_per_worker cannot be a negative number")
			return 0, err
		}
		return num, nil
	} else {
		err = errors.New("items_per_worker can only be integer type")
		return 0, err
	}
}

func getElement(elem []string, line int) (common.Element, error) {
	if len(elem) != 2 {
		err := errors.New("csv lines can only have two fields")
		return common.Element{}, err
	}

	name := elem[1]

	id, err := strconv.Atoi(elem[0])
	if err != nil {
		str := fmt.Sprintf("parsing error in line %d: \"%s\" is not a valid id, please use integers for id field",
			line, err.Error()[23:23+len(elem[0])])
		log.Println(str)
		err := errors.New(str)
		return common.Element{}, err
	}

	return common.Element{Key: id, Value: name}, nil
}

func getTask(element common.Element, filter func(common.Element, string) bool, parity string) *filterTask {
	return &filterTask{
		element: element,
		filter:  filter,
		parity:  parity,
	}
}

func getMaxWorkers() int {
	corenum := runtime.NumCPU()
	if corenum > 10 {
		return 10
	}

	return corenum
}
