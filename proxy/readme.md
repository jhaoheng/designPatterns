# proxy

- 當某個服務取得重複資訊過多, 造成服務壓力過大, 可透過 proxy 調配壓力
- ex:
    - 當 client 端 request 的次數越多, 則會直接跟後方的 service 進行互動
    - 當重複的 request 次數越多, 就會給後方 service 做越多相同的操作
    - 此時透過 proxy 來定義 lazying loading
    - 將相同的內容, 定義在 proxy 端, 不直接從 service 中取得資料
- ex:
    - 就像 nginx proxy 一樣, 作為一個 delegate
    - 前方有 nginx 負責處理判斷要送到 upstream 的哪一台 server

## 適用
- Lazy initialization
- Access control
- Local execution of a remote service.
- Loggin requests
- Caching request result.
- Smart reference

## 優點
- 可額外設計 proxy 的內容, 用來加強服務
- 透過 proxy, 可以自己管理 lifecycle, 不用管後面的 lib
- 透過 Open/Closed Principle, 可以無限建立新的 proxies 無需管理後面的 services

## 缺點
- 複雜性上升
- 可能會有延遲現象