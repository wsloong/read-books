// 练习6.1: 为bit数组实现下面这些方法
/*
func (*IntSet) Len() int // return the number of elements
func (*IntSet) Remove(x int) // remove x from the set
func (*IntSet) Clear() // remove all elements from the set
func (*IntSet) Copy() *IntSet // return a copy of the set
*/

package main

import "bytes"

import "fmt"

// IntSet 包含一组小的非负整数，它的零值表示空集合
type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, x%64
	return len(s.words) > word && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, x%64
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= (1 << bit)
}

func (s *IntSet) UnionWith(p *IntSet) {
	for i, word := range p.words {
		if i < len(s.words) {
			s.words[i] |= word
		} else {
			s.words = append(s.words, word)
		}

	}
}

func (s *IntSet) String() string {
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

func (s *IntSet) Len() int {
	var length int
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				length++
			}
		}
	}
	return length
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, x%64
	if s.Has(x) {
		s.words[word] ^= 1 << bit
	}
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	_s := &IntSet{words: []uint64{}}
	for _, word := range s.words {
		_s.words = append(_s.words, word)
	}
	return _s
}

func main() {
	var x IntSet
	x.Add(1)
	x.Add(3)
	x.Add(5)
	fmt.Println(&x, "==>", x.Len())

	x.Remove(5)
	fmt.Println(&x, "==>", x.Len())

	y := x.Copy()
	fmt.Println(&y, "==>", y.Len())

	x.Clear()
	fmt.Println(&x, "==>", x.Len())
}
