package adapter_object

/*
- Adaptee
*/

import "fmt"

type IThirdParty interface {
	DoSomething()
}

type ThirdParty struct{}

func NewThirdParty() IThirdParty {
	return &ThirdParty{}
}

func (tp *ThirdParty) DoSomething() {
	fmt.Println("3rd do something")
}
