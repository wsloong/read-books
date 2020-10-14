## Mutex 互斥锁

### 临界区

并发编程中，如果程序中的一部分会被并发访问或修改，为了避免并发访问导致意想不到的结果，
这部分程序需要保护起来，这部分被保护起来的程序，就叫做临界区

### Go的并发原语

1 Mutex: 互斥锁
2 RWMutex: 读写锁
3 WaitGroup: 并发编排
4 Cond: 条件变量
5 Channel: 通道

* 共享资源。并发地读写共享资源，会出现数据竞争(data race)的问题,需要 `Mutex`、`RWMutex`这样的并发原语来保护
* 任务编排。需要`goroutine`按照一定的规律执行，而`goroutine`之间有相互等待或者依赖的顺序，常使用`WaitGroup`或者`Channel`来实现
* 消息传递。消息交流已经不同`goroutine`之间的线程安全的数据交流，常使用`channel`

### Mutex的基本用法

`sync`包中定义了`Locker`接口，`Mutex`和`RWMetux`都实现了这个接口

```
type Locker interface {
    Lock()
    Unlock()
}
```
进入临界区要调用`Lock`方法，退出临界区调用`Unlock`方法

```
func main() {
	// 互斥锁保护计数器
	var mux sync.Mutex
	// 计数器的值
	var count int

	// 确保所有的goroutine都完成
	var wg sync.WaitGroup
	// 10个goroutine
	wg.Add(10)

	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 累加10万次
			for j := 0; j < 100000; j++ {
				// 使用互斥锁保护临界区
				mux.Lock()
				count++
				// 退出临界区
				mux.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
```

使用结构体对内部细节封装
```
func main() {
	// 封装好的计数器
	var counter Counter

	var wg sync.WaitGroup
	wg.Add(10)

	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 累加10万次
			for j := 0; j < 100000; j++ {
				counter.Incr() // 受到锁保护的方法
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.Count)
}

type Counter struct {
	counterType int
	Name        string

	// 一般把Mutex放在要控制的字段上面，然后使用空格把字段分隔开来
	mux   sync.Mutex
	count uint64
}

// Incr 内部加1的方法，受到锁的保护
func (c *Counter) Incr() {
	c.mux.Lock()
	c.count++
	c.mux.Unlock()
}

// 得到计数器的值，也需要锁的保护
func (c *Counter) Count() uint64 {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.count
}
```

### 问答
问：如果`Mutex`已经被一个`goroutine`获取了锁，其他`goroutine`只能等待；那么等这个锁释放后，等待中的`goroutine`中哪一个会优先获取到锁呢？

答：在正常状态下，所有等待锁的 `goroutine`按照 `FIFO` 顺序等待。唤醒的`goroutine`不会直接拥有锁，而是和新请求锁的`goroutine`竞争锁的拥有，新请求锁的`goroutine`具有优势:它正在CPU上执行，可能有好几个；所以刚唤醒的`goroutine`有很大可能会竞争失败，那么它会被放到等待队列的前面，如果一个等待的`goroutine`超过`1ms`没有获取到锁，它会将锁转变成`饥饿模式`。

`饥饿模式`下，锁的所有权将从`unlock`的`goroutine`直接交给等待队列中的第一个，新来的`goroutine`不会尝试获取锁，而是被放到等待队列的尾部
如果一个等待的`goroutine`获取了锁，并且满足一下其中的任何一个条件

* 它是队列中的最后一个
* 它等待的时间小于`1ms`

它会将锁的状态转换为正常状态

### 其他参考内容

* 查看汇编代码的命令: go tool compile -S file.go
* 查看`race`: 在编译、测试、运行时候增加 `-race`参数即可； `go run -race main.go`
* https://colobu.com/2018/12/18/dive-into-sync-mutex/

### Mutex的实现

```
type Mutex struct {
	state int32
	sema  uint32
}

```
state int32类型;被设计成一个复合型字段

mutexWaiters | mutexStarving | mutexWoken | mutexLocked

mutexLocked: 1bit位,标示锁是否被持有
mutexWoken: 1bit位, 唤醒标志
mutexStarving: 1bit位,饥饿标志
mutexWaiters: 剩余位; 一个先进先出的队列,保存等待的goroutine数量.



CAS: compare-and-swap; 原子操作;原子性保证这个指令总是基于最近的值进行计算,如果同时有其他线程已经修改了这个值,返回失败.

锁的`Unlock`可以被任意的`goroutine`调用释放锁,即使是没持有这个互斥锁的`goroutine`也可以这么操作;
这是因为,`Mutex`本身并没有包含持有这把锁的`goroutine`信息.所以`Unlock`也不会对此进行检查.
所以要遵循`谁申请,谁释放`的原则
