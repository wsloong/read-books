# 第九章 基于共享变量的并发

## 9.1 竞争条件

避免数据竞争：
* 不要去写变量(不可能)
* 避免从多个goroutine访问更新变量；确保只有一个goroutine可以访问和修改变量，其他goroutine使用一个channel来发送给指定goroutine来查询更新变量
    * 示例：`code/bank1`
    * 流水线中每个阶段都避免在将变量传送到下一个阶段时再去访问它，那么对这个而变量的所有访问就是线性的。变量会被绑定到流水线的一个阶段，传送完之后被绑定到下一个。这种称为串行绑定
* 互斥：允许多个goroutine去访问变量，但是同一时刻最多只有一个goroutine去访问
练习: `exercises/9.1`

## 9.2 sync.Mutex互斥锁

* `sync.Mutex`开箱即用
* 确保成对出现
* 使用`defer`语句，确保程序结束释放锁
* `Lock`和`Unlock`之间是临界区，可以随意读取和修改

使用二元信号量(一个容量为1的buffered channel)来控制同一时刻只有一个goroutine访问变量
示例:`code/bank2`，`code/bank3`

## 9.3 sync.RWMutex读写锁

* `Rlock` 只能在临界区共享变量没有写操作可用，如过有(比如更新缓存、计数加1)，请使用互斥锁。
* `RWMutex`只有当获得锁的大部分goroutine都是读操作,才能带来好处
* `RWMutex`需要更复杂的内部记录，所以会让它比一般的无竞争锁的mutex慢一些

## 9.4 内存同步

`code/bank3`中的`Balance`方法使用互斥锁的原因:
* 确保不会再其他操作(比如Withdraw)中间执行
* `同步`不仅仅是一堆goroutine执行顺序问题，同样涉及到内存的问题

每一个处理器都有`local cache`，为了效率，对内存的写入一般会在每一个处理器中缓冲，并在必要时候一起`flush`到主存，
这种情况下这些数据会以当初goroutine写入顺序不同，而以不同的顺序被提交到主存，
像`channel通信`或者`互斥锁`操作这样的原语会使处理器将其聚集的写入`flush`并`commit`，
这样`goroutine`在某个时间点上的执行结果才能被其它处理器上运行的`goroutine`得到。

## 9.5 sync.Once初始化

下面的代码片段使用了懒初始化。`不是并发安全的`

```
var icons map[string]image.Image

func loadIcons() {
    icons = map[string]image.Image{
        "spades.png": loadIcon("spades.png"),
        "hearts.png": loadIcon("hearts.png"),
        "diamonds.png": loadIcon("diamonds.png"),
        "clubs.png": loadIcon("clubs.png"),
    }
	// 上面的语句有一种可能会被重排成下面的
	// icons = make(map[string]image.Image)
	// icons["spades.png"] = loadIcon("spades.png")
	// icons["hearts.png"] = loadIcon("hearts.png")
	// icons["diamonds.png"] = loadIcon("diamonds.png")
	// icons["clubs.png"] = loadIcon("clubs.png")

	// 一个goroutine在检查icons是非空时候，也不能确定这个变量的初始化流程已经走完了
	// 可能只是塞了个空map，里面的值没有填充
}

// NOTE: not concurrency-safe!
func Icon(name string) image.Image {
    if icons == nil {
        loadIcons() // one-time initialization
    }
    return icons[name]
}
```

最简单切正确的保证所有goroutine能够观察到`loadIcons`效果的方式，是使用一个mutex来同步
```
var mu sync.Mutex // guards icons
var icons map[string]image.Image
// Concurrency-safe.
func Icon(name string) image.Image {
    mu.Lock()
    defer mu.Unlock()
    if icons == nil {
        loadIcons()
    }
    return icons[name]
}
```
代价就是没法对该变量进行并发访问，即使是该变量已经被初始化了且不会进行变动。可以使用允许多读的锁：
```
var mu sync.RWMutex // guards icons
var icons map[string]image.Image

// Concurrency-safe.
func Icon(name string) image.Image {
    // 首先获取一个写锁，查询map，然后释放锁
    mu.RLock()
    if icons != nil {
        icon := icons[name]
        mu.RUnlock()
        return icon
    }
    mu.RUnlock() // acquire an exclusive lock

    // 不释放共享锁的话，没有任何办法将一个共享锁升级为一个互斥锁。
    // 这里必须重新检查icons变量是否为nil，以防止在执行这一段代码的时候
    // icons变量已经被其它goroutine初始化过了
    mu.Lock()
    if icons == nil { // NOTE: must recheck for nil
        loadIcons()
    }
    icon := icons[name]
    mu.Unlock()
    return icon
}
```
上面的代码可以使用`sync.Once`来完美解决；`sync.Once`需要一个互斥锁`mutex`和一个`boolean`变量来记录初始化是否完成了。
```
var loadIconsOnce sync.Once
var icons map[string]image.Image

// Concurrency-safe.
func Icon(name string) image.Image {
    loadIconsOnce.Do(loadIcons)
    return icons[name]
}
```
练习: `exercises/9.2`

## 9.6 竞争条件检测
在`go build`，`go run`或者`go test`命令后面加上`-race`的flag.


