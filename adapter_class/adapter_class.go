package adapterclass

type IAdapter interface {
	IService
	ConvertToServiceFormat() interface{}
}

type Adapter struct {
	TP *ThirdParty
	S  *Service
}

func NewAdapter() IAdapter {
	adapter := &Adapter{
		TP: &ThirdParty{},
		S:  &Service{},
	}
	return adapter
}

func (a *Adapter) ConvertToServiceFormat() interface{} {
	return a.TP.StringToInt(a.S.Data)
}

func (a *Adapter) SetData(data string) {
	a.S.SetData(data)
}
