package taintflow

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Run(name string) ([]byte, error) {
	return archive(name)
}
