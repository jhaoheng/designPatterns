package adapter_object

import "fmt"

type IService interface {
	ConnectDB()
	SetData()
}

type Service struct {
	Adapter Adapter
	MyData  string
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ConnectDB() {
	fmt.Println("connection service")
}

func (s *Service) SetData() {
	s.MyData = "data"
}
