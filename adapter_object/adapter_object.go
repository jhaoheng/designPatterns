package adapter_object

import "fmt"

type Adapter struct {
	ThirdParty
}

func (a *Adapter) ConvertToJson() {
	fmt.Println("adapter convert to json data")
}
