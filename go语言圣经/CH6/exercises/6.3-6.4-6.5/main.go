// 练习 6.3： (*IntSet).UnionWith会用|操作符计算两个集合的交集，
// 我们再为IntSet实现另外的几个函数
// IntersectWith(交集：元素在A集合B集合均出现),
// DifferenceWith(差集：元素出现在A集合，未出现在B集合),
// SymmetricDifference(并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A)

// 练习6.4: 实现一个Elems方法，返回集合中的所有元素，用于做一些range之类的遍历操作

// 练习 6.5： 我们这章定义的IntSet里的每个字都是用的uint64类型，但是64位的数值可能在32位的平台上不高效。
// 修改程序，使其使用uint类型，这种类型对于32位平台来说更合适。
// 当然了，这里我们可以不用简单粗暴地除64，可以定义一个常量来决定是用32还是64，这里你可能会用到平台的自动判断的一个智能表达式：32 << (^uint(0) >> 63)
package main

import (
	"bytes"
	"fmt"
)

// 定义一个常量决定是32为还是64为平台
const BitNum = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

func (s *IntSet) Add(x int) {
	word, bit := x/BitNum, x%BitNum
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/BitNum, x%BitNum
	return len(s.words) > word && s.words[word]&(1<<uint(bit)) != 0
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] |= word
			continue
		}
		s.words = append(s.words, word)
	}
}

func (s *IntSet) Len() int {
	var length int
	for _, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < BitNum; j++ {
			if word&(1<<uint(j)) != 0 {
				length++
			}
		}
	}
	return length
}

func (s *IntSet) Remove(x int) {
	word, bit := x/BitNum, x%BitNum
	if s.Has(x) {
		s.words[word] ^= (1 << bit)
	}
}

func (s *IntSet) Copy() *IntSet {
	words := make([]uint, len(s.words))
	copy(words, s.words)
	return &IntSet{words}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < BitNum; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", BitNum*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// IntersectWith (交集：元素在A集合B集合均出现),
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, word := range t.words {
		if i >= len(s.words) {
			break
		}
		s.words[i] &= word
	}
}

// DifferenceWith (差集：元素出现在A集合，未出现在B集合),
func (s *IntSet) DifferenceWith(t *IntSet) {
	t1 := t.Copy()
	t1.IntersectWith(s)
	for i, word := range t1.words {
		if i < len(s.words) {
			s.words[i] ^= word
		}
	}
}

// SymmetricDifference(并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A)
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] ^= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

// Elems 方法，返回集合中的所有元素，用于做一些range之类的遍历操作
func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		for j := 0; j < BitNum; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, BitNum*i+j)
			}
		}
	}
	return elems
}

func main() {
	var x IntSet
	x.Add(1)
	x.Add(3)
	x.Add(5)
	x.Add(8)
	x.Add(10)
	fmt.Println(&x)

	var y IntSet
	y.Add(3)
	y.Add(6)
	y.Add(8)
	y.Add(9)
	fmt.Println(&y)

	// (&x).IntersectWith(&y)
	// fmt.Println(&x)
	// (&x).DifferenceWith(&y)
	// fmt.Println(&x)

	(&x).SymmetricDifference(&y)
	fmt.Println(&x)

}
