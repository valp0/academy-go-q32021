package services

type homeSvc struct{}

// Returns a homeSvc type.
func NewHomeSvc() homeSvc {
	return homeSvc{}
}

func (homeSvc) Inform() string {
	return "At this stage, available endpoints are /read, /fetch and /async."
}
