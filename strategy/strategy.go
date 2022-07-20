package strategy

import "fmt"

// type Sum struct {
// }

// func (s Sum) Execute(a int, b int) int {
// 	return a + b
// }

// 魔法書庫
type ICalculer interface {
	Execute(int, int) int
}

// 魔法師, 裝備 interface
type Calculer struct {
	Cal ICalculer
}

// 魔法師, 執行魔法
func (c *Calculer) Execute(a int, b int) int {
	fmt.Println("hello")
	return c.Cal.Execute(a, b)
}

// 第一個魔法
type Add struct {
}

func (add Add) Execute(a int, b int) int {
	return a + b
}

// 第二個魔法
type Minus struct {
}

func (m Minus) Execute(a int, b int) int {
	return a - b
}
