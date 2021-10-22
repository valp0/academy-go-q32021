package services

import "log"

type homeSvc struct{}

// Returns a homeSvc type.
func NewHomeSvc() homeSvc {
	return homeSvc{}
}

// Returns an informative string about the available endpoints.
func (homeSvc) Inform() string {
	msg := "At this stage, available endpoints are /read, /fetch and /async."
	log.Println(msg)
	return msg
}
