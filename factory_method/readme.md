# factory method
> https://github.com/tmrts/go-patterns/blob/master/creational/factory.md

## 目的
- 物件依據指定的參數建立後, 就具有該特性方法
	- ex: 儲存檔案到該參數指定的位置
	- ex: 發送 message 到該參數指定的對象
- 所以, 具有相同方法的物件, 集中在一起, 建立 interface, 就叫 factory
  - 雖然是不同的 struct, 但有相同的 method, 因此建立 interface
	- ex: 不同種類的汽車、儲存位置、發送 sns(ios/android)、通訊錄資料
- ex:
	- 儲存資料，可存放於 `記憶體`, `檔案`, `tmpfile`, `database` 這幾個儲存位置
	- 使用者 new(儲存類型), 就可以取得物件
	- 透過物件方法, 儲存資料內容與資料名稱
- 使用
	- pkg.NewNotification(pkg.FatherType).DoPhone()
	- pkg.NewNotification(pkg.FatherType).DoMessage("hello")
