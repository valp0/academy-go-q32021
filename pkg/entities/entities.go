package entities

type Element struct {
	Key   int         `json:"id"`
	Value interface{} `json:"name"`
}

type CSV struct {
	Path     string
	Body     [][]string
	Elements []Element
}
