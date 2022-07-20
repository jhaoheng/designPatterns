package adapter_object

import "fmt"

type IService interface {
	ConnectDB()
	SetXMLData()
	GetXMLData()
}

type Service struct {
	Adapter Adapter
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ConnectDB() {
	fmt.Println("connection service_1")
}

func (s *Service) GetXMLData() {
	fmt.Println("get xml data")
}

func (s *Service) SetXMLData() {
	fmt.Println("set xml data")
}
