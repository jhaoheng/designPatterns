package strategy

import (
	"fmt"
	"testing"
)

func Test_strategy(t *testing.T) {
	// 準備不同的魔法
	add := Calculer{Add{}}
	minus := Calculer{Minus{}}

	// 執行魔法
	fmt.Println(add.Execute(1, 2))
	fmt.Println(minus.Execute(1, 2))

	//
	// sum := Calculer{Sum{}}
	// sum.Execute(2, 3)
}
