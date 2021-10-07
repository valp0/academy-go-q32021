package common

// Element, file struct, etc

var FilePath string

type Element struct {
	Key   int    `json:"id"`
	Value string `json:"name"`
}

type Elements [][]Element

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
