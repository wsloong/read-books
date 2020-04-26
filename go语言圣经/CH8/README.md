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
示例`code/clock2`
练习`exercises/8.1`, `exercises/8.2`

## 8.3 示例：并发的Echo服务

服务端：`code/reverb1` => `code/reverb2`
客户端：`code/netcat2`

## 8.4 Channels

channels用于`goroutine`之间的通信，每个`channel`都有一个特殊的类型。比如`chan int`, `chan string`
可以使用make函数创建一个`channel`：`ch := make(chan int)`.
channel的零值也是nil；
和map类似，channle也对应一个make创建的底层数据结构的引用
复制和函数传递channle，只是拷贝了一个channle引用。
channle可以比较，只有2个channle引用的是相同的对象，才为真，channle也可以和nil比较
channel使用`<-`符号发送和接收数据
```
ch <- x     // send
x = <-ch    // receive
<-ch        // receive result is discarded
```
`close(ch)`可以关闭channel，关闭的channel不能发送数据，接收操作依然可以接受之前已经成功发送的数据，没有数据将产生一个零值的数据。
channel可以设置容量，称为带缓存的channel
```
ch = make(chan int)     // 非缓存channel
ch = make(chan int, 0)  // 非缓存channel
ch = make(chan int, 3)  // 缓存channel，容量为3
```

### 8.4.1 不带缓存的Channels

无缓存的channel发送操作将阻塞发送者的goroutine，直到另一个goroutine在相同channel上执行接收操作，反之依然。
基于无缓存的channel的发送和接收操作将导致两个goroutine做一次同步操作，所以无缓存的通道也被称为同步channel。
当通过一个无缓存Channels发送数据时，接收者收到数据发生在唤醒发送者goroutine之前。
在并发编程中，x事件发生在y事件之前，表示是要保证在此之前的事件都已经完成了；
我们x事件和y事件是并发的，不是意味着x和y就一定是同时发生的，只是不确定这两个事件发生的先后顺序。
示例：`code/netcat3`
练习：`exercises/8.3`

### 8.4.2 串联的Channels（Pipeline）

一个channel的输出可以作为下一个channel的输入，这就是所谓的管道(pipeline)
示例：`code/pipeline1`


对于通道是否关闭，可以通过`value, ok <-channel`，多返回一个ok，如果为true表示成功从channel接收到值，false表示channel已经被关闭了
更简洁的处理是使用`range`循环，当channel关闭，没有值可接收时候调出循环
示例：`code/pipeline2`
并不需要关闭每个通道，只要当需要告诉接收者goroutine，所有的数据已经全部发送时才需要关闭channel，
当一个channel没有被引用时候会被Go的垃圾回收器自动回收(这和打开的文件不同，打开的文件在不使用时候要调用Close方法)。
向一个已关闭的channel发送数据、重复关闭一个channel、关闭一个ni值的channel都会panic。

### 8.4.3 单方向的Channel

Go提供单方向channel类型，分别用于只发送(`chan<- T`)和只接收(`<-chan T`)的channel
双向类型的channel可以转换为单向类型的channel，但是单向类型不能转换为双向类型
示例`code/pipeline3`

### 8.4.4 带缓存的Channels
`ch = make(chan string 3)`创建3个容量的通道。
如果缓存channel满，发送操作会阻塞；如果缓存channel是空的，接收操作会阻塞。
`cap()`函数返回channel的容量。
下面代码，并发想三个站点发送请求，返回最先得到的值
```
func mirroredQuery() string {
    responses := make(chan string, 3)
    go func() { responses <- request("asia.gopl.io") }()
    go func() { responses <- request("europe.gopl.io") }()
    go func() { responses <- request("americas.gopl.io") }()
    return <-responses // return the quickest response
}
```
示例：`code/cake`

## 8.5 并发的循环

示例`code/thumbmail`(注意体会`makeThumbnails6`)
练习:`exercises/8.4`

