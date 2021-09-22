package utils

import (
	"encoding/json"
	"log"
)

// A function that returns a JSON containing what it receives
func JsonResponse(r interface{}) ([]byte, error) {

	// Returned JSON will be contained in a "data" field, and the
	// received interface will be contained in either a...
	type body struct {
		Data interface{} `json:"data"`
	}

	// ... "message" if it is a string or an...
	type message struct {
		Message interface{} `json:"message"`
	}

	// ... "items" field in any other case.
	type items struct {
		Items interface{} `json:"items"`
	}

	var resp body

	// We create our response body.
	_, isStr := r.(string)
	if isStr {
		resp = body{message{r}}
	} else {
		resp = body{items{r}}
	}

	// Parsing response struct to JSON.
	jRes, err := json.Marshal(resp)
	if err != nil {
		log.Println("Error converting message struct to JSON", err)
	}

	return jRes, err
}
