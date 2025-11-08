# 原因分析与解决方案
## 一、两处切片越界与并发问题
### 1. 闭包延迟绑定导致切片越界
程序中异步启动的 `goroutine` 里闭包的变量声明较慢，导致切片越界：
- 当 `len(consumeMSG) >= ConsumeNum` 时，主 `goroutine` 启动 `go func() { m := consumeMSG[:ConsumeNum]; fn(m) }()`；
- 但 `goroutine` 启动后不会立刻执行，主 `goroutine` 会继续往下走，执行到 `consumeMSG = consumeMSG[ConsumeNum:]`
- 等异步 `goroutine` 真正执行时，`consumeMSG` 可能已经长度较钝，甚至是空切片了，这时候执行 `consumeMSG[:ConsumeNum]`，就会触发 **切片越界 Panic**（`runtime error: slice bounds out of range`）。

### 2. 超时消费时切片越界
当（`len(consumeMSG) < ConsumeNum`），且超过 5 分钟未批量消费时，代码执行：
```go
m := consumeMSG[:ConsumeNum] // 比如 consumeMSG 只有 3 条，截取到 5，直接越界
```

### 3.并发安全问题
- 切片 `consumeMSG` 被主 `goroutine` （`append`、截取）和多个异步 `goroutine` （读取）同时操作，但**切片不是并发安全的**，可能导致数据竞争
- 即使没越界，异步 `goroutine` 里的 `m` 是 `consumeMSG` 的切片引用，主 `goroutine` 后续 `append` 可能**覆盖**底层数组数据


## 二、解决方案
对应上述问题一一解决，**上锁、不使用闭包、拷贝数据(防止覆盖)**
(具体修改结果见目录下同名go文件)
### 1. sync.Mutex
- 所有对 `consumeMSG` 的操作都要先 `mu.Lock()`，操作完再 `mu.Unlock()`
- 避免主 `goroutine` 写 `consumeMSG` 时，异步 `goroutine` 读，导致数据竞争。

### 2. copy
- `make`新切片并`copy`数据；
- 异步 `goroutine` 接收 `data []MSG` 作为参数，操作的是copy后的独立数据，和原 `consumeMSG` 无关

### 3. 超时消费按实际长度拷贝
- 用 `batchLen := len(consumeMSG)` 获取实际未消费数量，再创建对应长度的切片

