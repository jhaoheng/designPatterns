package adapter_object

/*
- Adaptee
*/

import "fmt"

type IThirdParty interface {
	DoSomething()
	SetData(data string)
}

type ThirdParty struct {
	Data string
}

func NewThirdParty() IThirdParty {
	return &ThirdParty{}
}

func (tp *ThirdParty) SetData(data string) {
	tp.Data = data
}

func (tp *ThirdParty) DoSomething() {
	fmt.Println(tp.Data)
}
