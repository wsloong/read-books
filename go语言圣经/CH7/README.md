# 第七章 接口

* 对其他类型行为的抽象和概括
* 满足隐式实现

## 7.1 接口的约定
接口类型是抽象类型
它不会暴露它所代表的对象的内部值的结构符合这个对象支持的基础操作的集合
它只会展示出他们自己的方法
示例`code/bytecounter`

## 7.2 接口类型
接口类型描述了一系列方法的结合，实现了这些方法的具体类型是这个接口的实例。
实例的方法顺序和接口的方法定义顺序没有影响
接口定义支持内嵌
```
type ReadWriter interface {
    Read(p []byte) (n int, err error)
    Writer
}
```

## 7.3 实现接口的条件
一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口。

对于具体类型T，它的一些接收者是类型T本身，也有T的指针，
一个T类型的参数调用一个T的方法是合法的，只要这个参数是一个变量，编译器会隐式的获取它的地址，
这仅仅是一个语法糖：T类型的值不拥有所有*T指针的方法，那么它可能只实现了更少的接口

```
type IntSet struct { /* ... */ }
func (*IntSet) String() string

// 无法通过编译，但是如果String()方法的接收者是一个值，而不是指针就没有任何问题
var _ = IntSet{}.String() // compile error: String requires *IntSet receiver

// 但是可以在一个IntSet值上调用
var s IntSet
var _ = s.String() 

// 由于只有指针类型的IntSet有String()方法，所以也只有指针类型的IntSet实现了fmt.Stringer接口
var _ fmt.Stringer = &s // OK
// 无法通过编译;
var _ fmt.Stringer = s // compile error: IntSet lacks String method
如果String()方法的接收者是值类型，那么无论该类型的值或这指针都实现了该接口
```

接口类型封装和隐藏具体类型和它的值，即使具体类型有其他的方法，也只有接口类型暴露出来的方法会被调用
```
var w io.Writer
w = os.Stdout
w.Write([]byte("hello")) // OK: io.Writer has Write method
w.Close() // compile error: io.Writer lacks Close method；即使os.Stdout有Close()方法也不能调用
```

空接口类型`interface{}`可以接收任意值的赋值