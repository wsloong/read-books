// bit数组的示例

// bit数组中一个成员可以存储64位,代表64个不同的数
// 将某一位置1即 代表将其加入数组中
// x/64获取的是元素在bit数组中的第几个成员,比如x=116 x/64=1 即元素可能在Bit[1]中
// x%64获取元素具体的位置x=116 x%64=52 则其应在Bit[1]的第52位上

package main

import (
	"bytes"
	"fmt"
)

// IntSet 包含一组小的非负整数，它的零值表示空集合
type IntSet struct {
	words []uint64 // 无符号的8byte
}

// Has 返回值x是否在集合中
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add 添加一个非零值到集合中
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)

	// for循环是为了扩展bit数组成员，比如原始有两个成员bit[0] bit[1]
	// 如果要加入元素155(大于64*2=128)原始集合无法标示(只有128位),需要在添加一个成员，初始化为0
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	// 增加元素时候，一位代表一个元素，如果要添加155，则将bit[2]中的第27位设置为1，也就是s.words[word] |= 1 << bit
	// 之所以使用|=是因为集合的成员已经存储了其他值，直接使用等号赋值会清空其他值，使用或等(|=)其他位置为1的标志位不会改变
	s.words[word] |= 1 << bit
}

// UnionWith 集合s和t的并集
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String 让集合按照"{1 2 3}"格式输出, 使用fmt包输出时候会自动调用该函数
// 相比原书的例子，这里我改成了使用值接受者，那么就可以不用调用显示调用s.String()
func (s IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	// 这里会自动调用x.String()方法
	// 如果String()方法定义时候是用指针接受者定义
	// 这里需要显示调用x.String()或者使用&x来调用
	fmt.Println(x) // {1 9 144}
	// fmt.Println(x.String()) // {1 9 144}

	y.Add(9)
	y.Add(42)
	fmt.Println(y) // {9 42}
	// fmt.Println(y.String()) // {9 42}

	x.UnionWith(&y)
	// 这里我改成了打印指针，看看显示效果
	// 可以看出编译器自动解指针调用了String()方法
	fmt.Println(&x)                   // {1 9 42 144}
	fmt.Println(x.Has(9), x.Has(123)) // true false
}
