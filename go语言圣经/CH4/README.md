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

## 4.3 Map

Map上一个`无序`的`key/value`对的集合，所有的`key`都是不同的；
Map通过`key`检索、更新或删除对应的`value`都是常数时间复杂度的；
Map中所有的key、value都有着相同的类型，key和value之间可以不同；
Map中`key`必须支持`==`比较运算符(浮点数不适合);
Map的元素通过`key`对应的下标语法访问和赋值；
删除Map中的元素使用`delete`函数：`delete(ages, "charlie")`;

Map的创建：
    * `make`函数：`ages := make(map[string]int)`
    * 字面量方式创建：`args := map[string]int{"alice": 31, "charlie": 34}`
    * 创建`nil`Map：`var args map[string]int`; 创建nil值的map`args`存入元素将会`panic`；但是查找、删除、len和range循环都是安全的；
```
	var m map[string]int        // nil值map
    fmt.Println(m == nil)       // true
    fmt.Println(len(m) == 0)    // true

    // 因为nil值的map没有引用任何哈希表，这里会没有任何打印
	for k, v := range m {
		fmt.Println(k, "=>", v) 
	}

    // 多返回一个值`ok`来标示key是否存在
	value, ok := m["hello"]
	fmt.Println(value, ok)      // 0 false
	delete(m, "hello")          // 不会panic

    // 重新初始化后才能向map存数据
    m = make(map[string]int)
	m["go"] = 1
	m["python"] = 2
	m["java"] = 3
	m["go"]++       // 
	fmt.Println(m)	// map[go:2 java:3 python:2]
```

不能对Map的元素取地址操作`_ = &ages["bob"]`：这是因为map可能随着数量的增长而重新分配内存空间，从而导致之前的地址无效；
Map遍历可以使用`range`(更常用)或者`for i := 0; i < len(m); i++ {...}`；
Map迭代顺序是不确定的，可以使用下面方法实现：
```
import "sort"

keys := make([]string, 0, len(ages))
for key := range ages {
    keys = append(keys, key)
}

// 排序
sort.String(keys)

for _, key := range keys {
    fmt.Printf("%s\t%d\n", key, ages[key])
}
```

和slice一样，map之间不能进行相等比较，唯一例外是和`nil`进行比起我叫。
要判断俩个map是否包含相同的key、value。可以通过一个循环实现
```
func equal(x, y map[string]int) bool {
    if len(x) != len(y) {
        return false
    }
    for k, xv := range x {
        if yv, ok := y[k]; !ok || yv != xv {
            return false
        }
    }
    return true
}
```

示例: `code/dedup`(实现类似set功能)

虽然Map的`key`必须是可比较的，但是通过方法可以不可比较的`key`类型；下面使用`slice`作为key的一些处理示例：
```
var m = make(map[string]int)

func k(list []string) string { return fmt.Sprintf("%q", list) }

func add(list []string) { m[k(list)]++ }

func Count(list []string) int { return m[k(list)] }
```

示例:`code/charcount`(统计输入中每个Unicode码点出现的次数)

示例:`code/graph`(value也可以是一个聚合类型，比附map或者slice)

## 4.4 结构体

结构体是一种聚合的数据类型，由零个或多个任意类型的成员聚合成的实体；
成员可以通过`.`操作符访问或赋值；
结构体可以对成员取地址，然后通过指针访问或赋值；
结构体成员通过大小写绝对是否可以导出；
一个结构体`S`不能包含自身，但是可以包含`*S`指针类型的成员，可以创建递归的数据结构，比如链表和树结构；
```
type Employee struct {
    ID            int
    Name, Address string    // 相邻成员类型可以合并到一行
    DoB           time.Time
    Position      string
    Salary        int
    ManagerID     int
}

// 明了一个Employee类型的变量dilbert;成员值为类型零值
var dilbert Employee

// 通过点操作对成员赋值
dilbert.Salary += 5000

// 点操作配合指针一起工作
var employeeOfTheMonth *Employee = &dilbert
employeeOfTheMonth.Position += " (proactive team player)"
// 和上面的语句效果一样
(*employeeOfTheMonth).ManagerID = 10
```

示例：`code/treesort`

结构体类型的零值是每个成员都是零值，通常会将零值作为最合理的默认值；`bytes.Buffer`、`sync.Mutex`都是开箱即用的。
空结构体(没有任何成员)，写作`struct{}`，大小为0。不包含任何信息。

### 4.4.1 结构体字面量

结构体字面值可以指定每个成员的值；
1，按照结构体成员定义的顺序为每个结构体成员制定一个字面值(不推荐这种，如果后续增加成员还要考虑顺序问题)
```
type Point struct{ X, Y int }
p := Point{1, 2}
```
2，以成员名称和相应的值来初始化，可以包含部分或全部成员
```
type Point struct{ X, Y int }
p := Point{X:1, Y:2}
```

结构体可以作为函数的参数和返回值，较大的结构体考虑效率的话，通常会用指针方式传入和返回；
如果在函数内部修改结构体成员，必须使用指针；
Go语言中，所有的函数参数都是值拷贝传入，指针传入也是拷贝指针的值；

### 4.4.2 结构体比较

如果结构体全部成员都是可比较的，那么结构体也是可比较的；
可比较的结构体类型和其他可比较的类型一样，可以用于map的key；
```
type address struct {
    hostname string
    port     int
}

hits := make(map[address]int)
hits[address{"golang.org", 443}]++
```

### 4.4.3 结构体嵌入和匿名成员
Go语言有一个特性让我们只声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫匿名成员。
匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针;

示例：`code/embed`

匿名成员也要一个隐式的名字，因此不能同时包含两个类型相同的匿名成员；
匿名成员也有可见性规则约束(大写可以导出，小写不能导出);

## JSON
JavaScript对象表示法（JSON）是一种用于发送和接收结构化信息的标准协议;
`encoding/json`对于JSON的编码和解码都有良好的支持；
Go语言将结构体转JSON的过程叫编组(marshaling)，使用`json.Marshal()`完成；
只有导出的结构体成员才能被编码；
`Tag`关联到该成员的元信息字符串；`omitempty`选项标示当Go语言结构体成员为空或零值是不生成JSON对象；
编码的逆操作是解码，对应将JSON数据解码为Go语言的数据结构，Go语言中一般叫`unmarshaling`，通过`json.Unmarshal`函数完成

示例：`code/movie`
