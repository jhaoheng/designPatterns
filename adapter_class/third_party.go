package adapterclass

import "strconv"

type IThirdParty interface {
	StringToInt(data string) int
}

type ThirdParty struct {
}

func (tp *ThirdParty) StringToInt(data string) int {
	i, _ := strconv.Atoi(data)
	return i
}
