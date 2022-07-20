# 什麼是 `Strategy Design Pattern`
- https://skyyen999.gitbooks.io/-study-design-pattern-in-java/content/strategy.html
- https://medium.com/@felipedutratine/solid-open-closed-principle-in-golang-3d23e8258a45

# readme
- 目的: 將各種可以互換的演算法(策略)包裝成一個類別
	- 包裝成一個類別是重點, 如果沒有包裝成一個類別, 就有點像是 Decorator
- 好處
	- 透過 pkg 設定不同的彈藥包, 執行的方式相同
		- 因為彈藥包有自己的類別, 所以已經被定型, 也不能隨意改動 = open-closed principle
		- 但可以增加自己的彈藥包, 只要滿足相同 struct 的設計
	- 可以有效簡化 pkg 的使用方式
- 步驟
	- 每個不同的物件, 具有相同的 method
	- 組合成 interface
	- 建立一個`入口物件`, 繼承 interface
	- 建立`入口物件`的 method, 跟每個物件的 method 相同