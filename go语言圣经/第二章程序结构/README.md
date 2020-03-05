## 程序结构

### 2.1命名
* 必须有一个字母或者下划线开头，后面可以跟任意数量的字母、数字或下划线
* 大小写字母是不同的变量(区分大小写)
* 关键字不能作为自定义名字，下面是25个关键字

|        |         |        |           |             |
| ---    |  ---    |    --- |   ---     |   ---       |
| break  | default | func   | interface | select      |
| defer  | go      | map    | struct    | else        |
| goto   | package | switch | const     | fallthrough | 
| if     | range   | type   | continue  | for         |
| import | return  | var    | case      | chan        |

* 预定义的名字(对应内建的常量、类型和函数)。这些不是关键字，可以再定义中重新使用(避免过度引起语义混乱)
```
内建常量: true false iota nil

内建类型: int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
float32 float64 complex128 complex64
bool byte rune string error

内建函数: make len cap new append copy close delete
complex real imag panic recover
```
* 变量名尽量短小，见名知意
* 推荐使用`驼峰式`命名(parseRequestLine)

### 2.2声明

* Go 程序对应一个或多个以`.go`为文件后缀的源文件中
* 每个源文件以bao的声明语句开始(`package main`)
* `import`可以导入依赖的其他包
* 然后是包一级的类型、变量、常量、函数的声明语句

| 关键字 | 类型 |
|  ---   | ---  |
| var    | 变量 |
| const  | 常量 |
| type   | 声明一种类型 |
| func   | 函数 |

示例`code/boiling/main.go`, `code/ftoc/main.go`

### 2.3变量

`var 变量名称 类型 = 表达式`其中 `类型`或者`= 表达式`2个部分可以省略其中的一个。

    * 省略`类型`；根据初始值表达式推导变量的类型
    * 省略`= 表达式`；使用类型零值初始化该变量； 
    * 类型零值：数值->0、布尔->false、字符串->""、接口或引用(slice/map/chan/函数)->nil
    *数组或结构体等聚合类型对应的零值是每个元素或字段都是对应该类型的零值

```
var i, j, k int // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string
var f, err = os.Open(name)  // os.Open returns a file and an error
```

#### 2.3.1简短变量声明

`名字 := 表达式`; 变量类型根据表达式来自动推导

* `:=`是一个变量声明语句，`=`是一个变量赋值操作
* 简短变量声明也可以用函数的返回值来声明和初始化变量
    ```
        f, err := os.Open(name)
    ```
* 在`相同的语法域`声明过了，简短变量声明对这些已经声明过的变量只有赋值行为
    ```
        in, err := os.Open(inFile)      // 声明了in和err两个变量
        // ...
        out, err := os.Open(outFile)    // 只声明了out一个变量，然后对已声明的err进行赋值操作
    ```
* 
* 简短变量声明语句中至少要声明一个新的变量
* 简短变量声明只有对已经在`同级语法域`声明的变量才和赋值操作语句等价，如果是`外部语法域`，那么会在当前语法域重新声明一个新变量
* 支持多个赋值
```
i := 100
a, b := true, 1
```

#### 2.3.2指针

* 一个变量对应一个保存了变量对应类型值的内存空间
* 一个指针的值是另一个变量的地址(内存中的存储位置)
* 不是每个值都会有一个内存地址，但是对于每一个变量必然有对应的内存地址
* 可以直接读或者更新对应变量的值，而不需要知道变量的名字
* 变量有时候被称为可寻址的值
* 任何类型的指针零值都是nil
* 指针之间可以进行相等比较，只有指向同一个变量或者都是nil才相等
* 对于聚合类型的每个成员(比如结构体的每个字段、数组的每个元素)也都对应一个变量，可被取地址

`var x int`,`&x`将产生一个指向该整数变量的指针(`p`)，指针对应数据类型`*int`;
`指针p`保存了`x`变量的内存地址，`*p`表达式对应`p指针`指向的变量的值
```
x := 1
p := &x             // *int类型，执行x
fmt.Println(*p)     // "1"
*p = 2              // 等价于 x = 2
fmt.Println(x)      // "2"
```
示例 `code/echo4/main.go`

#### 2.3.3 new函数

`new(T)`将创建一个T类型的匿名变量，初始化为T类型的零值，返回返回变量地址，返回的指针类型为`*T`;

#### 2.3.4 变量的生命周期

变量的生命周期：程序运行期间变量有效存在的时间间隔。
* 包一级变量，生命周期和整个程序的运行周期是一致的
* 局部变量: 从创建到该变量不再备引用为止，变量的存储空间可能被回收

编译器会自动选择在栈还是堆上分配局部变量的存储空间
```
var global *int

func f() {
    var x int
    x = 1
    global = &x
}

func g() {
    y := new(int)
    &y = 1
}
```
`f函数`中的`x`变量必须在堆上分配，因为它在函数退出后依然可以通过包一级的`global`变量找到；这个`x`局部变量从函数f中逃逸了。