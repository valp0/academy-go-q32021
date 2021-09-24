package handlers

import "net/http"

type homeHandler struct{}
type readHandler struct{}

type IHomeHandler interface {
	Home(w http.ResponseWriter, r *http.Request)
}

type IReadHandler interface {
	ReadLocalCsv(w http.ResponseWriter, r *http.Request)
}

func NewHomeHandler() IHomeHandler {
	return homeHandler{}
}

func NewReadHandler() IReadHandler {
	return readHandler{}
}
