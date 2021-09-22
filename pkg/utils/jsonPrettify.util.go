package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// A function that prettifies JSON strings.
func Prettify(ugly []byte) string {
	var prettified bytes.Buffer
	err := json.Indent(&prettified, ugly, "", "  ")
	if err != nil {
		fmt.Println("Couldn't prettify received JSON string:", err)
		return string(ugly) // Will return received string in case of error.
	}

	return prettified.String()
}
