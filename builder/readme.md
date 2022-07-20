# builder
> https://github.com/tmrts/go-patterns/blob/master/creational/builder.md

## 目的

- 當使用者想使用某個物件, 可以自行根據該物件的專屬配件, 裝飾該物件的外觀
- ex: 
    - 一台車可以有
		- 不同的輪子
		- 不同的引擎
		- 不同的顏色
	- 但一台車子的骨架相同
	- 所以我們定義車子的骨架, 同時把輪子、引擎、顏色當作配件
	- 我們會定義配件的 type, const
	- 定義這台骨架的 interface 中包含哪些方法

## 設計前提
- 需要知道物件的主體
- 需要知道物件的參數有哪些

## 如果要使用 builder pattern

1. 找出骨架是什麼
2. 找出骨架的配件有哪些
3. 定義出每個配件的類型 (ex: type Something string)
	- 依據類型定義該配件的固定參數有哪些 (const variables, 給予一個可供人類閱讀的字串)
	- 定義出該類型的 method
4. 定義出骨架的 interface, 其中包含會用到的 method

## 程式設計

- 定義該物件的結構名稱 (struct)
- 定義物件的內部, 可用於歸納的參數型態
	- ex: type color string, 不會定義成 type red string 或 type blue string 等
	- 設定該參數型態下, 會用到的 const variables, 並且給予可閱讀的名稱
- 依照參數型態, 在物件結構下, 建立對應的 method
	- method 的輸入參數, 需要強迫指定已設定好的參數類型
	- 回傳該物件的 point struct
- 設定該物件 struct 中的參數 (依據剛剛建立好的參數型態)
- 建立該物件的 new method
- 建立該物件的 interface
- 設定該物件的 new method, 回傳 interface
