package repo

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/valp0/academy-go-q32021/pkg/entities"
)

type CSV entities.CSV

func NewCsvService(path string) (CSV, error) {
	f := CSV{Path: path}
	err := f.Read()
	if err != nil {
		return CSV{}, err
	}

	err = f.FillElements()
	if err != nil {
		return CSV{}, err
	}

	return f, nil
}

// Reads a file content into its body field.
func (f *CSV) Read() error {
	raw, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer raw.Close()

	csvReader := csv.NewReader(raw)
	file, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	f.Body = file
	return nil
}

// Fills elements field from raw Body field.
func (f *CSV) FillElements() error {
	elements := []entities.Element{}

	for i, elem := range f.Body {
		name := elem[1]

		id, err := strconv.Atoi(elem[0])
		if err != nil {
			str := fmt.Sprintf("Parsing error in line %d: “%s” is not a valid ID, please use integers for ID field.",
				i+1, err.Error()[23:23+len(elem[0])])
			fmt.Println(str)
			newErr := errors.New(str)
			return newErr
		}

		elements = append(elements, entities.Element{Key: id, Value: name})
	}

	f.Elements = elements
	return nil
}

// Returns the the element matching the provided ID, or an error if no element matches ID.
func (f *CSV) QueryById(id string) ([]entities.Element, error) {
	pokeId, err := strconv.Atoi(id)
	if err != nil {
		err := errors.New("“" + id + "” is not a valid ID, only integer type is allowed as an ID.")
		return []entities.Element{}, err
	}

	for _, e := range f.Elements {
		if e.Key == pokeId {
			return []entities.Element{e}, nil
		}
	}

	err = errors.New("No element with ID “" + id + "” was found.")
	return []entities.Element{}, err
}
