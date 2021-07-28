package main

import (
	"decorator/pkg"
	"fmt"
	"log"
)

/*
## 目的
- 在不修改原本物件內部的程式碼, 動態的延伸物件功能
- Decorator 提供彈性的作法來延伸物件的功能

## 經驗
- 不像 Adapter pattern, 物件是透過 injection 來 decorated
- Decorators 不應該修改物件的 interface
*/

func main() {
	DecorateWithStcuct()
}

//
func DecorateWithStcuct() {
	//
	pizza := &pkg.VeggeMania{}
	//add cheese
	pizzaWithCheese := &pkg.CheeseTopping{
		Pizza: pizza,
	}
	// add tomato topping
	pizzaWithCheeseAndTomato := &pkg.TomatoTopping{
		Pizza: pizzaWithCheese,
	}
	fmt.Println(pizzaWithCheeseAndTomato.GetPrice())
}

//
func DecorateWithMethod() {
	f2 := LogDecorate(pkg.LogPrint)
	f2("I am MAX")
}

//
type Object func(string) string

func LogDecorate(fn Object) Object {
	return func(msg string) string {
		log.Println("Starting, ", msg)

		result := fn(msg)

		log.Println("Execution is completed with the result", result)

		return result
	}
}
