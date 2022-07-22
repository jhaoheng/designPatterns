# readme

- https://refactoring.guru/design-patterns/facade
- Facade is a structural design pattern that provides a simplified interface to a library, a framework, or any other complex set of classes.


```
Facades hide a complex or extensive API behind a small interface you can interact with without understanding all the details behind the scenes. Like we mentioned before, go-cloud is an excellent example of both adapter and facade, as it hides the complex details of talking to cloud providers behind a straightforward interface.
```

## 優點
- 在複雜的子系統中, 可以將之獨立設計在 Facade 中, 透過簡單的介面使用

## 缺點
- 過於龐大的 Facade Lib, 會讓整包變成 a god object, 耦合性可能會過重