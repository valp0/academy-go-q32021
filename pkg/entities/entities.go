package entities

type Element struct {
	Key   int         `json:"id"`
	Value interface{} `json:"name"`
}

type file interface {
	Read() error
	// Write() (int, error)
}

type CSV struct {
	Path     string
	Body     [][]string
	Elements []Element
}
