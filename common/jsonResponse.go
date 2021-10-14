package common

// Return what is given as an argument inside a JSON response

import (
	"encoding/json"
	"log"
)

// A function that returns a JSON containing what it receives.
func JsonResponse(r interface{}) []byte {
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
		log.Println("Error converting response struct to JSON.", err)
		return []byte{}
	}

	return jRes
}
