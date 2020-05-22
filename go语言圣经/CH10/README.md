# 第10章 包和工具

## 10.5 包的匿名导入
* 导入一个不使用的包会导致一个编译错误
* `_`可以匿名导入一个包，这样会计算包级变量的初始化表达式和执行包的`init`初始化函数
示例：`code/jpeg`, 练习：`exerciese/10.1`

## 10.工具

```
go doc time
go doc time.Since

godoc -http :8000

go list 命令可以查询可用包的信息

go list github.com/go-sql-driver/mysql
go list ...
go list gopl.io/ch3/...
go list -json hash

命令行参数 -f 则允许用户使用text/template包（§4.6）的模板语言定义输出文本的格式。下
面的命令将打印strconv包的依赖的包，然后用join模板函数将结果链接为一行，连接时每个结
果之间用一个空格分隔：
go list -f '{{join .Deps " "}}' strconv
go list -f '{{.ImportPath}} -> {{join .Imports " "}}' compress/...
```
