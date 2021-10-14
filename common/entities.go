package common

// Element, response struct, etc

// Represents a key-value pair.
type Element struct {
	Key   int    `json:"id"`
	Value string `json:"name"`
}

/*---------------------------------------------------------*/
// jsonResponse.go
type body struct {
	Data interface{} `json:"data"`
}

type message struct {
	Message interface{} `json:"message"`
}

type items struct {
	Items interface{} `json:"items"`
}

/*---------------------------------------------------------*/
