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