package repo

import (
	"encoding/csv"
	"fmt"
	"os"
)

// A function that receives a path to a csv file,
// and returns a 2D string slice with its content.
func ReadCsv(path string) ([][]string, error) {

	// Opening csv file and handling error.
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Unable to open file", path, "\nReason:", err)
		return nil, err
	}
	defer f.Close()

	// Reading from csv file and handling error.
	csvReader := csv.NewReader(f)
	file, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Unable to read file", path, "\nReason:", err)
		return nil, err
	}

	return file, nil
}
