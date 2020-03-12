package main

import "testing"

type m struct {
	s1       string
	s2       string
	outorder bool
}

var tables = []m{
	m{"123", "1", false},
	m{"123", "123", false},
	m{"123", "321", true},
	m{"123", "321", true},
	m{"中国", "中", false},
	m{"中国", "中国", false},
	m{"中国", "国中", true},
}

func TestIsOutOrder(t *testing.T) {
	for index, table := range tables {
		if table.outorder != IsOutOrder(table.s1, table.s2) {
			t.Fatalf("index:%d, s1:%s, s2:%s, outorder:%t", index, table.s1, table.s2, table.outorder)
		}
	}

}
