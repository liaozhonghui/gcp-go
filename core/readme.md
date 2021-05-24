## GO 语言 Core 核心

1. GO 语言工作区和程序实体
2. GO 语言数据类型

- 数组，切片， container，字典 map
- channal 通道
- struct interface
- func
- interface 类型的使用
- 指针类型的使用，与可寻址的类型与数据

3. GO 语言程序执行

- if else switch select
- gorutine
- defer panic

4. GO 语言实战

### GO 语言编程模式

- 切片 slice, 接口 interface, 时间 time 库，性能:编写 go 程序的规范
- 错误处理 if err != nil， 使用 callback 进行错误传递
- Functional options 模式 ...fns 进行 pipeline 化， 处理重载和 Builder 的情况
- 嵌入和委托，反转控制，不是由控制逻辑依赖业务逻辑，而是由业务逻辑依赖控制逻辑
- Map-Reduce: Map, Reduce, Filter, 应对需要统计的业务场景。（需要处理成泛型编程模式）
- GO 语言代码生成，typeAssert=>.(type)进行转型操作，Reflection 反射，函数模板（go generate 注释生成），第三方工具，genny, gen
- decorator， 泛型处理
- Pipeline, 使用 decorator 函数处理，或者使用 channel 进行处理
- Vistior 模式：将算法与操作对象的结构分离的一种方法，遵循开放封闭原则。通过 Builder 模式，将数据转化为一系列的资源，最后用 Visitor 模式来进行迭代处理这些 Reources.
