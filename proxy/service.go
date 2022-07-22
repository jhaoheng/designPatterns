package proxy

import "fmt"

type IService interface {
	Operation()
}

type Service struct{}

func NewService() IService {
	return &Service{}
}

func (s *Service) Operation() {
	fmt.Println("service operation success")
}
