# 第11章 测试

Go语言的测试依赖一个`go test`测试命令和一组按照约定方式编写的测试函数；

## 11.1 go test

* `_test.go`后缀的源文件在执行`build`时候不会被构建，它们是`go test`测试的一部分
* 在`_test.go`中，以`Test`为函数前缀的为测试函数，`Benchmark`为函数前缀的为基准测试函数，`Example`为函数前缀的是示例函数
* `go test`会遍历所有的`*_test.go`文件中符合上述命名规则的函数，生成一个`临时的mai包`用于调用响应的测试函数，接着构建并运行、报告测试结果，最后清理测试中生成的临时文件
* `go test -v `每个测试函数的名字和运行时间
* `go test -run `对应一个正则表达式，只有被匹配到的测试函数才会被运行，比如`go test -v -run="French|Canal"`

## 11.2 测试函数

* 测试函数必须以`Test`开头。可选的后缀名必须以大写字母开头；`TestName(t *testing.T)`
* 测试函数的签名如下
```
func TestName(t *testing.T) {}
```
示例：`code/word1`, `code/word2`, `exerciese/11.1`

## 11.2.1 随机测试
随机测试如何知道希望的输出结果？
* 编写另一个对照函数，使用简单和清晰的算法，针对相同的随机输入检查两者的输出结果
* 生成的随机输入的数据遵循特定的模式
* [go-fuzz](https://github.com/dvyukov/go-fuzz)


第二种方式的示例: `code/random_palindrome`

## 11.2.2. 测试一个命令
示例： `code/echo`;
测试代码和产品代码在同一个包，虽然是main包，也对应的`main`入口函数，但是在测试时候`main`包只是`TestEcho`测试函数导入的一个普通包，
里面`main`函数并没有被导出，而是被忽略了

## 11.2.3. 白盒测试
黑盒测试只需要测试包公开的文档和API行为
白盒测试有访问包内部函数和数据结构的权限(之前的`echo`程序，更新了out包级变量)
示例：`code/storage`

## 11.3 测试覆盖率

* `测试能证明缺陷存在，而无法证明没有缺陷`
* 覆盖率是指在测试中至少被运行一次的代码占总代码数的比例

示例：`CH7/code/eval` 

运行示例 使用`go tool`命令，该命令在`$GOROOT/pkg/tool/${GOOS}_${GOARCH}`。
* `go test -run=Coverage -coverprofile=c.out ./eval`
    * `用 go test -cover` 显示摘要
    * `-covermode=count` 将在每个代码块插入一个计数器而不是布尔标志量，在统计结果中记录了每个块的执行次数，可衡量哪些是被频繁执行的热点代码。
* `go tool cover -html=c.out`

## 11.4 基准测试
* 以`Benchmark`为前缀
* 参数固定有一个 `*testing.B` 类型的参数
* 默认情况下`go test`不运行基准测试，`-bench=` 命令行标志参数可以手工指定要运行的基准测试函数，该参数值是一个正则表达式
* `-bench=.` 可以匹配所有的基准测试函数

## 11.5 剖析

收集数据
```
go test -cpuprofile=cpu.out           // CPU剖析数据标识了最耗CPU时间的函数
go test blockprofile=block.out        // 阻塞剖析则记录阻塞goroutine最久的操作,例如系统调用、管道发送和接收，还有获取锁等
go test -memprofile=mem.out           // 堆剖析则标识了最耗内存的语句
```
使用`go tool pprof`命令

## 11.6. 示例函数

以 `Example` 为函数名开头

```
func ExampleIsPalindrome() {
fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
fmt.Println(IsPalindrome("palindrome"))
// Output:
// true
// false
}
```
示例代码的好处
* 示例代码是真实Go代码，需要接受编译器的编译时候检查；作为文档:`godoc`会将`ExampleIsPalindrome`作为`IsPalindrome`文档的一部分展示
* `go test`执行测试的时候也会运行示例代码，如果示例中有`// Output:`格式的注释，测试工具会检查示例代码的输出与注释是否匹配
* 提供真实的演练场
