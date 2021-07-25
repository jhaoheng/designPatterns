package main

import (
	"fmt"
	"sync"
)

/*
# 目的
限制物件實例化時是單一物件, 唯一物件, 通常用於全局 (全局變數), ex: 連線
只有這個物件, 可以處理指定事件

# Rule of thumb
- 因為是全局的資源, 大多數情況下會降低可測試性
*/

func main() {
	s := NewSingleton()
	s["this"] = "that"

	s2 := NewSingleton()
	fmt.Println(s2["this"])
}

type Singleton map[string]string

var (
	once     sync.Once
	instance Singleton
)

func NewSingleton() Singleton {
	once.Do(func() {
		instance = make(Singleton)
	})
	fmt.Println("NewSingleton instantiation")
	return instance
}
