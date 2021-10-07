package repo

import (
	"encoding/csv"
	"os"
	"strconv"
)

func addPokemon(id int, name string) error {
	f, err := os.OpenFile(os.Getenv("PATH"), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	alreadyStored, err := alreadyStored(id)
	if err != nil {
		return err
	}

	if alreadyStored {
		return nil
	}

	csvWriter := csv.NewWriter(f)
	pokemon := []string{strconv.Itoa(id), name}

	err = csvWriter.Write(pokemon)
	if err != nil {
		return err
	}

	csvWriter.Flush()
	return nil
}

func alreadyStored(id int) (bool, error) {
	f, err := os.Open(os.Getenv("PATH"))
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
