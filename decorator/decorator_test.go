package decorator

import (
	"fmt"
	"testing"
)

/*
## 目的
- 在不修改原本物件內部的程式碼, 動態的延伸物件功能
- Decorator 提供彈性的作法來延伸物件的功能

## 經驗
- 不像 Adapter pattern, 物件是透過 injection 來 decorated
- Decorators 不應該修改物件的 interface
*/
//
func Test_pizza(t *testing.T) {
	//
	pizza := NewBasicPizze()

	//add cheese
	pizzaWithCheese := &CheeseTopping{
		Pizza: pizza,
	}

	// add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		Pizza: pizzaWithCheese,
	}
	fmt.Println(pizzaWithCheeseAndTomato.GetPrice())
}
