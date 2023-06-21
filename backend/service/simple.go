package service

type SimpleService interface {
	GetText() string
}

type SimpleServiceImpl struct{}

func (s *SimpleServiceImpl) GetText() string {
	return "test"
}

func NewSimpleService() SimpleService {
	return &SimpleServiceImpl{}
}
