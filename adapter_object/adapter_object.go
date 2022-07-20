package adapter_object

import "fmt"

type Adapter struct {
	Adapee ThirdParty
}

func (a *Adapter) ConvertToJson(data string) {
	data = fmt.Sprintf("flavored %v", data)
	a.Adapee.SetData(data)
}
