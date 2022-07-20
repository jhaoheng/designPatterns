# decorator

## 目的
- 在不修改原本物件內部的程式碼, 動態的延伸物件功能
  - 有點類似 open-closed principle
- Decorator 提供彈性的作法來延伸物件的功能

- 有點類似
	1. 建立一個物件類型 
		- 使用 type 建立, 物件不限於 struct, 也可以是 method
	2. 建立一個 mehtod (decorator), 以物件類型當作 input, 然後同時輸出物件類型
	3. 不更改該物件的內部, 但延伸該物件的功能

- ex:
	1. 有一個物件 struct, 有自己的 method
	2. 現在用 decorate 將 物件 當作輸入, 輸出也是該物件

## Rules of Thumb
- 不像 Adapter pattern, 物件是透過 injection 來 decorated
- Decorators 不應該修改物件的 interface
