package main

import "testing"

type t struct {
	src  string
	dest string
}

var tables = []t{
	t{"a", "a"},
	t{"a.go", "a"},
	t{"a/b/c.go", "c"},
	t{"a/b.c.go", "b.c"},
}

func TestBasename1(t *testing.T) {
	for _, table := range tables {
		if rest := Basename1(table.src); rest != table.dest {
			t.Fatalf("src:%s, expect:%s, get:%s\n", table.src, table.dest, rest)
		}
	}
}

func TestBasename2(t *testing.T) {
	for _, table := range tables {
		if rest := Basename2(table.src); rest != table.dest {
			t.Fatalf("src:%s, expect:%s, get:%s\n", table.src, table.dest, rest)
		}
	}
}
