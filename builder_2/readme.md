# Define

- Builder : 為創建一個Product對象的各個部件指定抽象接口。
- ConcreteBuilder : 
    - 實現 Builder 的接口以構造和裝配該產品的各個部件。
    - 定義並明確它所創建的表示。
    - 提供一個檢索產品的接口
- Director
    - 構造一個使用 Builder 接口的對象。
- Product
    - 表示被構造的複雜對象。
    - ConcreateBuilder 創建該產品的內部表示並定義它的裝配過程。
    - 包含定義組成部件的類，包括將這些部件裝配成最終產品的接口。

# Test

`go test builder.go builder_test.go -v -cover`

```
=== RUN   TestDog
--- PASS: TestDog (0.00s)
=== RUN   TestCat
--- PASS: TestCat (0.00s)
PASS
coverage: 65.0% of statements
ok      command-line-arguments  0.005s  coverage: 65.0% of statements
```