package repo

import (
	"encoding/csv"
	"os"

	"github.com/valp0/academy-go-q32021/common"
)

// Return pokemon element matching given ID
// Respond with prettified JSON (jsonResponse.go common)

// Returns the the element matching the provided ID.
func GetById(id string) (string, error) {
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

	elements, err := filterElement(id, content)
	if err != nil {
		return "", err
	}

	jRes := common.JsonResponse(elements)
	prettyRes := common.PrettifyJson(jRes)

	return prettyRes, nil
}
