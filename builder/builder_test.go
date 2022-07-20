package builder

import (
	"testing"
)

func Test_builder(t *testing.T) {
	/*
		- 設計前提
			- 需要知道物件的主體
			- 需要知道物件的參數有哪些
		- 程式設計
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
	*/
	NewCar().SetColor(RedColor).SetLight(Funture).SetWheel(RaceWheel).Build()
	NewCar().Build()
	//
	NewNotification().SetTitle(InfoTitle).SetBody(InfoBody).Send()
}
