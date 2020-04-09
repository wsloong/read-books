# 第五章 方法

## 6.1 方法的声明
在函数声明时，在其名字之前放上一个变量，既是一个方法
示例`code/geometry`

## 6.2 基于指针对象的方法
当方法的接受者变量本身比较大时，可以使用指针而不是对象来声明方法
```
func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}
```
约定：如果`Point`这个类有一个指针作为接收器的方法，那么所有`Point`的方法都必须有一个指针接收器。
如果`p`是`Point`类型的变量，但是方法需要一个`*Point`指针作为接受者，我们依然可以通过`p.ScaleBy(2)`来调用，
编译器会隐式的帮我们用`&p`去调用`ScaleBy`这个方法；

总结：
* 不管你的`method`的`receiver`是指针类型还是非指针类型，都是可以通过指针/非指针类型进行调用的，编译器会帮你做类型转换。
* `method`的`receiver`是指针还是非指针类型，需要考虑两方面
    * 对象本身是不是很大。如果很大，使用非指针变量，调用会产生一次拷贝
    * 如果用指针作为`receiver`，这个指针指向的始终是一块内存地址，就算你对其进行了拷贝

### 6.2.1 Nil也是一个合法的接收器类型
函数允许nil指针作为参数，那么类似的，方法也可以用nil指针作为接收器
下面的int链表例子中，nil代表空链表

```
// An IntList is a linked list of integers.
// A nil *IntList represents the empty list.
type IntList struct {
    Value int
    Tail *IntList
}
// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
    if list == nil {
        return 0
    }
    return list.Value + list.Tail.Sum()
}
```

## 6.3 通过嵌入结构体来扩展类型
示例`code/coloredpoint`

## 6.4 方法值和方法表达式

当你根据一个变量来决定调用同一个类型的哪个函数时，方法表达式就显得很有用了。

```
type Point struct{ X, Y float64 }
func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point
func (path Path) TranslateBy(offset Point, add bool) {
    var op func(p, q Point) Point
    if add {
        op = Point.Add
    } else {
        op = Point.Sub
    }

    for i := range path {
        // Call either path[i].Add(offset) or path[i].Sub(offset).
        path[i] = op(path[i], offset)
    }
}
```

## 6.5 示例: Bit数组

Go语言的集合一般使用`map[T]boo`这种形式来表示。
下面使用bit数组更好的表示它

一个bit数组通常会用一个无符号数或者称之为`字`的slice来表示，每一个元素的每一位都表示集合里的一个值
当集合的第i位被设置时，我们才说这个集合包含元素i
示例`code/intset`

## 6.6 封装

封装的好处：
* 因为调用方不能直接修改对象的变量值，其只需要关注少量的语句并且只要弄懂少量变量的可能的值即可
* 隐藏实现的细节
* 是阻止了外部调用方对对象内部的值任意地进行修改