# 第三章 基础数据类型

Go语言的四类数据结构：基础类型、复合类型、引用类型和接口类型
基础类型:
    数字、字符串、布尔型
复合类型
    数组、结构体
引用类型
    指针、切片、字典、函数、通道

## 3.1 整型

Go语言提供了有符号和无符号类型的整数运算

| | | | |
|-|-|-|-|
|8bit|16bit|32bit|64bit|
|int8|int16|int32(rune)|int64|
|uint8(byte)|uint16|uint32|uint64|

`Unicode字符`,`rune类型`是和int32等价的类型，通常用于标示一个Unicode码点；这两个名称可以互换使用。
`byte`和uint8类型是等价类型为那个，byte类型用于强调数值是一个原始的数据而不是一个小的整数。
`uintptr`是无符号的整数类型，没有指定具体大小，但是足以容纳指针。

* 算数运算符：

`+`、`-`、`*`、`/`、可以适用于整数、浮点数和复数;
取模运算符`%`仅适用于整数间运算；在Go语言中，`%`取模运算符的符号和被取模数的符号总是一致的，`-5%3`和`-5%-3`的结果都是`-2`；
除法运算符`/`的行为则依赖与操作数是否全部为整数，`5.0/4.0 == 5.0/4 == 5/4.0 = 1.25`、但是`5/4=1`；

算术运算符要小心结果溢出，超出的高位bit为将会被丢弃，如果原是数值是有符号的，而且最左边的bit位是1，结果可能是负的
```

// 无符号的uint8，范围[0 ~ 2^8-1]
// 有符号的uint8，范围[-2^7 ~ 2^7-1]
var u uint8 = 255
fmt.Println(u, u+1, u*u)    // 255, 0, 1

var i int8 = 127
fmt.Println(i, i+1, i*i)    // 127, -128, 1
```

* 比较运算符:

`==`、`!=`、`>`、`>=`、`<`、`<=`；比较运算符返回布尔类型。

* 位运算符
`&`：与运算
对应位同时为1才是1，否则为0；例如 00000100 & 00001111 => 00000100

`|`：或运算
对应位只要有一个为1就是1；例如 00000100 & 00001111 => 00001111；

`^`: 
* 一元运算符: 按位取反；例如 ^00010100 => 11101011
* 二元运算符: 异或，对应位相同为0，不同为1；例如 00010100 ^ 00001111 = 00011011

`&^`：位清空
`a &^ b`: 将b中为1的位 对应于a的位清0， a中其他位不变；例如 00010100 &^ 00001111 = 00010000

* `>>`右移和`<<`左移
左移：右边空出的位用 0 填补，高位左移溢出则舍弃该高位; `x<<n`等价于`2^n`；
右移: 无符号的用 0 填充左边空位；有符号的用符号位填，正数用 0，负数用 1

* 一些知识
```
func main() {
	var i uint8 = 20
	fmt.Println(^i, ^20)      // 235 -21
}

原因是 一个是有符号的数，一个是无符号的数
20默认是int类型，最高位是符号位，符号位取反得到的就是复数；

负数的二进制数怎么表示？
负数的二进制数是它对应的正数按位取反得到反码，再加上1得到的补码
例如
3的二进制：     00000000 00000000 00000000 00000011
反码：          11111111 11111111 11111111 11111100
补码：反码加1： 11111111 11111111 11111111 11111101
故-3的二进制为  11111111 11111111 11111111 11111101
```

不同整数类型之间可以通过`T(x)`显示将一个值从一个类型转化为另一种类型。但是对于将一个大的整数类
型转为一个小的整数类型，或者是将一个浮点数转为整数，可能会改变数值或丢失精度；

Go语言也提供了无符号数的运算，即使数值本身不可能出现负数，我们还是倾向于使用有符号的`int`类型，就像数组的长度:
```
medals := []string{"gold", "silver", "bronze"}
for i := len(medals) - 1; i >= 0; i-- {
fmt.Println(medals[i]) // "bronze", "silver", "gold"
}
```
如果`len()`函数返回一个无符号数，`i`也是无符号uint类型，那么`i >= 0`则永远为真，当`i == 0`时候，`i--`语句不会是-1，而是`uint`类型的最大值，此时`medals[i]`数组将会越界，程序会panic。

字符面值通过一队`单引号`直接包含对应字符，如ASCII中的'a'写法的字符面值；但是可以通过专一的数值来标示任意的Unicode码点对应的字符；
```
ascii := 'a'
unicode := '国'
newline := '\n'
fmt.Printf("%d %[1]c %[1]q\n", ascii)   // 97   a   'a'
fmt.Printf("%d %[1]c %[1]q\n", unicode) // 22269 国 '国'
fmt.Printf("%d %[1]q\n", newline)       // 10 '\n'
```

## 3.2 浮点数

Go语言提供了两种精度的浮点数：float32和float64；最大值对应`math.MaxFloat32`和`math.MaxFloat64`；
通常应该`优先使用float64类型`，它提供了更高的精度；
很小很大的数最好用科学计数法书写
```
const Avogadro = 6.02214129e23  // 阿伏伽德罗常数
const Planck = 6.62606957e-34 // 普朗克常数
```
可以用`fmt.Printf`函数的`%g`、`%e`、`%f`的形式来打印浮点数；

示例`code/surface/main.go`

## 3.3 复数

Go语言提供了两种精度的复数：complex64和complex128，分别对应float32和float6两种浮点数精度；
`complex`函数用于构建复数，`real`和`image`分别返回复数的实部和虚部
```
var x complex128 = complex(1, 2)    // 1 + 2i 等价于 x := 1 + 2i
var y complex128 = complex(3, 4)    // 3 + 4i
fmt.Println(x*y)                    // (-5+10i)
fmt.Println(real(x*y))              // -5
fmt.Println(imag(x*y))              // 10
```
如果一个浮点数面值或者十进制整数面值后面跟一个i，比如`2i`或者`3.1415i`;它将构成一个复数的虚部，实部为0；
复数可以进行比较，只有两个复数的实部和虚部都相等时候他们才是相等的；
`math/cmplex`包提供了处理复数的函数；

示例`code/mandelbrot/main.go`

## 3.4 布尔型

布尔型只有两个值：true和false。一元运算符`!`对应逻辑非操作；!true==false; !false==true
布尔值可以和`&&`和`||`操作符结合，且有短路行为；
`&&`的优先级比`||`高；

## 3.5 字符串

字符串是不可改变的字节序列；通常被解释为采用UTF8编码的Unicode码点(rune)序列；
`len`函数返回一个字符串中的`字节数据`，不是字符数目；
索引操作`s[i]`返回第i个字节， i必须是 `0 <= i < len(s)`;
`s[i:j]`基于原始s字符串的第i到j个字节(不包含j)生成一个新字符串；
`i`或者`j`都可能被忽略，当备忽略时候采用`0`作为开始位置，`len(s)`作为结束位置，比如`s[:]`， `s[:2]`, `s[3:]`;
`+`操作符(+=)可以将两个字符串连接成一个新的字符串；
字符串可以比较，通过逐个字节比较完成，比较的结果是字符串自然编码的顺序；

字符串是不可修改的，不变性意味着如果两个字符串共享相同的底层数据的话也是安全的，这使得复制任何长度的字符串代价是低廉的；
一个字符串`s`和对应的子字符串切片`s[7:]`的操作也可以安全的共享相同的内存，因此字符串切片操作也是低廉的；
这两种情况下都没有必要分配新的内存

### 3.5.1 字符串面值

字符串值也可以用字符串面值方式编写，只要将一系列字节序列包含在双引号即可："hello world"；
Go语言的源文件总是UTF8编码， 并且Go语言的文本字符串也以UTF8编码的方式处理；
一个原生的字符串面值形式是`...`，使用反引号代替双引号，在原生的字符串面值中，没有转义操作，可以跨行；

### 3.5.2 Unicode

早期，计算机世界只有一个ASCII字符集：美国信息交换标准代码。
ASCII，更准确地说是美国的ASCII，使用7bit(一个字节的低7位就可以表示，最高位是0)来表示128个字符：
包含英文字母的大小写、数字、各种标点符号和设备控制符。
随着计算机的普及，ASCII无法包含各种语言(中文、日文等等)的丰富的文本数据，就出现了Unicode编码；
一个Unicode码点的数据类型是int32，也就是Go语言中rune对应的类型；每个Unicode码点都是用同样的大小的32bit来标示；
这种方法虽然简单统一，但是相比ASCII码(只需要8bit或1字节)会浪费很多存储空间。

### 3.5.3 UTF-8

UTF8是一个将Unicode码点编码为字节序列的变长编码，由于是变长，就比完全使用Unicode(4个字节)占用更少的存储。
UTF8编码使用1到4个字节来表示每个Unicode码点，ASCII部分字符只使用1个字节，常用字符部分使用2或3个字节表示。
每个符号编码后第一个字节的高端bit位用于表示总共有多少编码个字节。
如果第一个字节的高端bit为0，则表示对应7bit的ASCII字符，ASCII字符每个字符依然是一个字节，和传统的ASCII编码兼容。
如果第一个字节的高端bit是110，则说明需要2个字节；
如果第一个字节的高端bit是1110，则说明需要3个字节；
如果第一个字节的高端bit是11110，则说明需要4个字节；
后续的每个高端bit都以10开头。更大的Unicode码点也是采用类似的策略处理。
```
0xxxxxxx                             runes 0-127    (ASCII)
110xxxxx 10xxxxxx                    128-2047       (values <128 unused)
1110xxxx 10xxxxxx 10xxxxxx           2048-65535     (values <2048 unused)
11110xxx 10xxxxxx 10xxxxxx 10xxxxxx  65536-0x10ffff (other values unused)
```
Go语言的源文件采用UTF8编码，并且Go语言处理UTF8编码的文本也很出色。
unicode包提供了诸多处理rune字符相关功能的函数（比如区分字母和数组，或者是字母的大写和小写转换等);
unicode/utf8包则提供了用于rune字符序列的UTF8编码和解码的功能。

```
import "unicode/utf8"

func main() {
	s := "hello, 世界"

    // 字符串包含13个字节，以UTF8形式编码，但是只对应9个Unicode字符
	fmt.Println(len(s))                         // 13
	fmt.Println(utf8.RuneCountInString(s))      // 9

	for i := 0; i < len(s); {
        // UTF8解码器
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}
```

每一次调用DecodeRuneInString函数都返回一个r和长度，r对应字符本身，长度对应r采用UTF8编码后的编码字节数目。
长度可以用于更新第i个字符在字符串中的字节索引位置。
但是这种编码方式是笨拙的，我们需要更简洁的语法。
幸运的是，Go语言的range循环在处理字符串的时候，会自动隐式解码UTF8字符串。
下面的循环运行如图3.5所示；需要注意的是对于非ASCII，索引更新的步长将超过1个字节。
![avatar](ch3-05.png)

```
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
```
无论显示调用`utf8.DecodeRuneInString`解码还是在range循环中隐式的解码，如果遇到一个错误的UTF8编码输入，将生成一个特别的Unicode字符`\uFFFD`,在印刷中这个符号通常是一个黑色六角或钻石形状，里面包含一个白色的问号"�".

UTF8字符串作为交换格式非常方便，但是程序内部采用rune序列可能更方便，因为rune大小一直，支持数组索引和方便切割；
将[]rune类型转换应用到UTF8编码的字符串，将返回编码的Unicode码点序列
```
str := "中国，世界"
fmt.Printf("% x \n", str) 	// e4 b8 ad e5 9b bd ef bc 8c e4 b8 96 e7 95 8c
r := []rune(str)
fmt.Printf("%x\n", r) 		// [4e2d 56fd ff0c 4e16 754c]
```
如果是将一个[]rune类型的Unicode字符slice或数组转为string，则对它们进行UTF8编码：
```
fmt.Println(string(r))	// 中国，世界
```

将一个整数转换为字符串意思是生成以包含对应Unicode码点字符的UTF8字符串，
如果对应码点的字符是无效的，则用\uFFFD无效字符作为替换：
```
fmt.Println(string(65))     // "A", not "65"
fmt.Println(string(1234567)) // "�"
```

### 3.5.4 字符串和Byte切片

标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。
strings包提供了许多如字符串的查询、替换、比较、截断、拆分和合并等功能。
bytes包也提供了很多类似功能的函数，但是针对和字符串有着相同结构的[]byte类型。
strconv包提供了布尔型、整型数、浮点数和对应字符串的相互转换，还提供了双引号转义相关的转换。
unicode包提供了IsDigit、IsLetter、IsUpper和IsLower等类似功能，它们用于给字符分类。

示例`code/basename`, `code/comma`

个字符串是包含的只读字节数组，一旦创建，是不可变的。相比之下，一个字节slice的元素则可以自由地修改。
字符串和字节slice之间可以相互转换
```
s := "abc"
b := []byte(s)
s2 := string(b)
```

bytes包和strings包同时提供了许多使用使用函数：
strings包的六个函数
```
func Contains(s, substr string) bool
func Count(s, sep string) int
func Fields(s string) []string
func HasPrefix(s, prefix string) bool
func Index(s, sep string) int
func Join(a []string, sep string) string
```
bytes包中也对应的六个函数
```
func Contains(b, subslice []byte) bool
func Count(s, sep []byte) int
func Fields(s []byte) [][]byte
func HasPrefix(s, prefix []byte) bool
func Index(s, sep []byte) int
func Join(s [][]byte, sep []byte) []byte
```

bytes包还提供了开箱即用的Buffer类型，用于字节slice的缓存。
当向`bytes.Buffer`添加任意字符串的UTF8编码时，最好使用`bytes.Buffer的WriteRune`
示例："code/printints/main.go"

### 3.5.5 字符串和数字的转换

除了字符串、字符、字节之间的转换，字符串和数值之间的转换也比较常见，有`strconv`包提供这类转换功能。

整数转字符串：
```
x := 123
y := fmt.Sprintf("%d", x)       // %b、%d、%o、%x对应二进制、十进制、八进制、十六进制的格式化
fmt.Println(y, strconv.Itoa(x))     // 123, 123

`FormatInt`和`FormatUint`函数可以用不同的进制来格式化数字
fmt.Println(strconv.FormatInt(int64(x), 2))     // 1111011
```

字符串转整数：
```
x, err := strconv.Atoi("123")             // x is an int
y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
```

有时候也会使用fmt.Scanf来解析输入的字符串和数字，特别是当字符串和数字混合在一行的时候，它可以灵活处理不完整或不规则的输入。
