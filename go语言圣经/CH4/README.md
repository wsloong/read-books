# 第四章 复合数据类型

复合数据类型，是以不同方式组合基本类型构造出来的数据类型；主要讨论四种类型：数组、slice、map和结构体。

数组和结构体是聚合类型，它们的值有许多元素或成员字段的值组成；
数组是有同构的元素组成：每个数组元素都是完全相同的类型；结构体则是异构的元素组成；

数组和结构体都是有固定内存大小的数据结构；
slice和map则是动态的数据结构，它们将根据需要动态增加。

## 4.1 数组

数组是有`固定长度`的`相同元素`组成的序列，一个数组可以有零个或多个元素组成。
数组可以通过下标访问元素，索引小标范围0 ～ len(array)-1;
默认情况下，数组每个元素都被初始化为元素类型对应的零值；
```
var r [3]int = [3]int{1, 2}
fmt.Println(r[2]) // 第二个元素没有指定，默认是int类型的零值：0      

// 这里定义了一个包含100个元素的数组r2,最后一个元素被初始化为-1，其他为元素类型零值：0
r2 := [...]int{99: -1}

// 数组字面量中，如果长度位置出现了`...`省略号，标示数组长度根据初始化值的个数来计算；
q := [...]int{1, 2, 3}
fmt.Printf("%T\n", q)       // [3]int
```

长度是数组类型的一个组成部分，因此`[3]int`和`[4]ing`是两种不同的数组类型，不能相互赋值和比较；
数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定。

数组也可以制定一个索引和对应值列表的方式初始化：
```
type Currency int

const (
    USD Currency = iota // 美元
    EUR                 // 欧元
    GBP                 // 英镑
    RMB                 // 人民币
)

symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

fmt.Println(RMB, symbol[RMB]) // 3 ￥
```

示例： `code/sha256/main.go`

数组作为函数实参时候，Go语言默认是传递的数组的副本，而不是数组的指针；
虽然可以显式的传递一个数组的指针，但是数组类型长度的限制，使用数组依然不方便。
除非明确的知道要处理特定大小数组(比如示例sha256)。大部分会使用slice代替数组；

## 4.2 切片

切片： 变长(没有固定长度)；每个元素都相同； `[]T`

创建切片
    * 字面量声明：`s1 := []int{0, 1, 2, 3, 4, 5}`、`var s []int`
    * 可以用`make`函数创建一个指定元素类型、长度和容量的slice，`make([]T, len, cap)`，容量可以省略，这时候容量等于长度
    * slice的切片操作; `s1 := s[i:j]` 返回一个新的slice，和原slice共享底层数组
    * 数组切片操作，`array[i:j]`

切片有三部分构成：指针、长度、容量
    * 指针：指向`第一个slice元素`对应的`底层数组元素的地址`
    * 长度：slice中的元素数量，len(s)
    * 容量：从slice开始位置到底层数组的结尾位置，cap(s)

使用`len(s) == 0`判断slice是否为空

示例： `code/rev/main.go`

slice不能进行比较，数组可以；
slice不直接支持比较运算符是因为
    * slice元素是间接引用的，一个slice甚至可以包含自身
    * slice不同时刻可能包含不同元素，因为底层数组的元素可能被修改

特例:
    * `[]byte`类型slice，可以通过`bytes.Equal`函数进行比较是否相等；
    * slice的零值是`nil`，因此slice可以和nil比较;

其他类型slice可以自定义比较函数，比如
```
func equal(x, y []string) bool {
    if len(x) != len(y) {
        return false
    }
    for i := range x {
        if x[i] != y[i] {
            return false
        }
    }
    return true
}
```

### ４.2.1 append函数

内置的append函数用于向slice追加元素
```
var runes []rune
for _, r := range "Hello, 世界" {
    runes = append(runes, r)
}
fmt.Printf("%q\n", runes) // ['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']
```
示例：`code/append`

### 4.2.2 Slice内存技巧

示例：`code/nonempty`
slice模拟stack：`code/slicestack`

### 关于copy函数
`copy(dst, src)`,将src切片复制到dst切片，替换dst相应位置元素；如果两个数组切片不一样大，按照较小的个数进行复制

```
s1 := []int{1, 2, 3, 4, 5} 
s2 := []int{5, 4, 3} 

copy(s2, s1) // 只会复制 s1 的前3个元素到 s2 中 
copy(s1, s2) // 只会复制 s2 的3个元素到 s1 的前3个位置
```
