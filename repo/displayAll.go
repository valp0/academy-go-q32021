package repo

import (
	"encoding/csv"
	"os"

	"github.com/valp0/academy-go-q32021/common"
)

// Return all pokemon elements in csv file
// Respond with prettified JSON (jsonResponse.go common)

// Returns a prettified JSON string of all the elements inside the csv.
func DisplayAll() (string, error) {
	path := os.Getenv("PATH")

	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	content, err := csvReader.ReadAll()
	if err != nil {
		return "", err
	}

	elements, err := getElements(content)
	if err != nil {
		return "", err
	}

	jRes := common.JsonResponse(elements)
	prettyRes := common.PrettifyJson(jRes)

	return prettyRes, nil
}
