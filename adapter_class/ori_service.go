package adapterclass

type IService interface {
	SetData(data string)
}

type Service struct {
	Data string
}

func (s *Service) SetData(data string) {
	s.Data = data
}
