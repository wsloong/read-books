# 第八章 Goroutines和Channels

Go语言中并发实现
* goroutine和channel,支持"顺序通信进程"(communicating sequential processes),简称CSP
* 多线程共享内存

## 8.1 Goroutines

当一个程序启动时候，主函数在一个单独的goroutine中运行，即`main goroutine`；
新goroutine会用go语句来创建；`go 函数/方法`
```
f()     // 调用f();会等待函数返回
go f()  // 创建一个goroutine执行f();不会阻塞
```
示例`code/spinner`

## 8.2 示例：并发的Clock服务

示例`code/clock1`。 可以使用`nc`或者`telnet`命令来模拟客户端，使用：`nc localhost 8000`;
示例`code/netcat1` 模拟客户端；如果同时开2个去链接`code/clock1`的服务端，可以看到服务器同时只能处理一个链接，
第二个必须等待第一个客户端完成工作，会没有输出。可以使用`go handleConn`,为每一个链接创建一个goroutine，见如下：
示例`code/