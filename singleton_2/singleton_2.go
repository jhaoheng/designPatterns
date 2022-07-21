package singleton2

/*
- 使用 sync.Once 來確保只會執行一次
*/

import (
	"fmt"
	"sync"
)

var once sync.Once

func GetDBConnection() {
	once.Do(func() {
		fmt.Println("Do connection to DB")
	})
}
