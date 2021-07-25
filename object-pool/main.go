package main

import "fmt"

/*
目的: 根據預期的需求, 先期準備與保留多個實例, 以備不時之需
ex: 預留的記憶體空間、預留的陣列大小、預留的連線數量、預留的 cache 狀態資訊

優點:
- 從資源成本上考量, 若 object 的初始化 > 維運, 則可以考慮 object-pool
- 如果需求並非穩定需求, 則維護 object-pool 資源的開銷會大於利益
- 由於物件被預先初始化, 會對性能有正面的影響
*/

func main() {
	p := NewObject(2)

	select {
	case obj := <-p:
		obj.Do()
		p <- obj
	default:
	}
}

type Object struct {
}

type Pool chan *Object

func NewObject(total int) Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- &Object{}
	}

	return p
}

func (o *Object) Do() {
	fmt.Printf("%v\n", o)
}
