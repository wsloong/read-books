# 第五章 函数

## 5.1 函数声明

函数包含`函数名`、`形参列表`、`返回值列表`(可省略)、`函数体`；
函数的类型被称为函数的标识符。
函数调用必须按照顺序提供实参，Go语言没有默认值；

## 5.2 递归

函数直接或间接调用自身
示例`code/findlisks1`
示例`code/outline`
```
以上示例结合CH1的`fetch`功能

go build -o fetch ../CH1/code/fetch/main.go 
go build -o findlinks1 ./code/findlisks1/main.go 
go build -o outline ./code/outline/main.go
./fetch https://golang.google.cn/ | ./findlinks1
./fetch https://golang.google.cn/ | ./outline
```

## 5.3 多返回值

Go语言函数支持返回多个值;调用者必须显式的将这些值分配给变量;如不使用某个值，可以使用`_`接收；
一个函数将所有的返回值都显示的变量名，那么该函数的return语句可以省略操作数。这称之为`bare return`
```
func CountWordsAndImages(url string) (words, images int, err error) {
    words, images, err = countWordsAndImages(doc)
    return              // 这里的return 相当于 return words, images, err
}

```
`bare return`会让代码难以理解，不宜过度使用

## 5.4 错误

内置的error是接口类型。error类型可能是nil(成功)或者non-nil(失败)。

### 5.4.1 错误处理策略

* 传播错误:错误信息链式组合，避免大写和换行符；错误信息描述要详尽；被调用函数会将调用信息和参数信息作为错误的上下文放到错误信息中并返回给调用者
* 重新尝试:限制重试的时间间隔和次数，防止无限制的重试
* 输出错误信息并结束程序，该策略只用在`main`中执行。对于库函数应该向上传播错误
* 有时候只需要输出错误信息就足够了，不需要中断程序。多用于debug
* 直接忽略错误

### 5.4.2 文件结尾错误(EOF)

io.EOF:这是个固定错误；文件结束引起的读取失败，只需要简单的比较就可以检测出这个错误

## 5.5 函数值

在Go中，函数被看着第一类值(first-class values)，也叫第一公民；
函数像其他值一样，拥有类型，可以被赋值给其他变量，传递给函数，从函数返回；
对函数值调用类似函数调用；
函数类型的零值是nil；因此函数值可以与nil比较；调用nil函数会引发panic；
示例`code/outline2`

## 5.6 匿名函数

匿名函数:`func`关键字后没有函数名

示例`code/squares`(闭包)
   * `squares`可以看到，变量的声明周期不由它的作用域决定:`squares`返回后，变量`x`仍然隐式的存在`f`中

示例 `code/toposort`(拓扑排序)
示例 `code/links`(广度优先)

### 5.6.1 警告：捕获迭代变量

```
var rmdirs []func()
for _, d := range tempDirs() {
    dir := d // NOTE: necessary!
    os.MkdirAll(dir, 0755) // creates parent directories too
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir)
    })
}

// ...do some work…
for _, rmdir := range rmdirs {
    rmdir() // clean up
}
```
代码中引入了一个临时变量`dir`；这是因为函数值中记录的是循环变量的内存地址，而不是循环变量某一时刻的值。以`d`为例，后续的迭代会不断更新`d`的值，
当删除操作执行时，for循环已完成，`d`中存储的值等于最后一次迭代的值。索引将`d`赋值给一个临时变量`dir`来解决这个问题

## 5.7 可变参数
参数数量可变的函数称为为可变参数函数。可变参数在最后一个参数类型前面加上`...`;
示例`code/sum`

## 5.8 Deferred函数
示例`code/title1`
示例中为了确保关闭了网络链接，`resp.Body.close`被调用了多次；`defer`延迟执行可以让这个操作变得简单(return,panic都会执行);具体查看`code/title2`
多条`defer`语句，执行顺序与声明顺序相反；

`defer`机制也常被用于记录何时进入和退出函数。示例`code/trace`

在循环体中的defer语句需要特别注意，因为只有在函数执行完毕后，这些被延迟的函数才会执行。
下面的代码会导致系统的文件描述符耗尽，因为在所有文件都被处理之前，没有文件会被关闭；
```
for _, filename := range filenames {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close() // NOTE: risky; could run out of file

    // ...process f…
}
```
解决方法是将循环体中的defer语句移至另外一个函数。在每次循环时，调用这个函数；
```
for _, filename := range filenames {
    if err := doFile(filename); err != nil {
        return err
    }}

func doFile(filename string) error {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()
    // ...process f…
}
```

## 5.9 Panic异常
`Panic`大多是程序运行时候产生，比如数组访问越界，空指针等；也可以调用`panic()`函数触发。
示例`code/defer2`

## 5.10 Recover捕获异常
在`defer`函数中调用了内置函数`recover`，并且定义该`defer`语句的函数发生了panic异常，recover会使程序从panic中恢复，并返回panic value。
导致panic异常的函数不会继续运行，但能正常返回。在未发生panic时调用recover，recover会返回nil。
```
func Parse(input string) (s *Syntax, err error) {
    defer func() {
        if p := recover(); p != nil {
            err = fmt.Errorf("internal error: %v", p)
        }
    }()
    // ...parser...
}
```