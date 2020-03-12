package main

import "testing"

type m struct {
	src  string
	desc string
}

var tables = []m{
	m{"123", "123"},
	m{"1234", "1,234"},
	m{"1234.56", "1,234.56"},
	m{"1234567890", "1,234,567,890"},
	m{"1234567.890", "1,234,567.890"},
	m{"-1234.56", "-1,234.56"},
	m{"+123456.78", "+123,456.78"},
	m{"-123", "-123"},
	m{"+123.23", "+123.23"},
}

func TestComma(t *testing.T) {
	for _, table := range tables {
		if res := Comma(table.src); res != table.desc {
			t.Fatalf("comma:src:%s, expect:%s, result:%s\n", table.src, table.desc, res)
		}
	}
}

func TestCommaWithoutRecursive(t *testing.T) {
	for _, table := range tables {
		if res := CommaWithoutRecursive(table.src); res != table.desc {
			t.Fatalf("commaWithoutRecursive: src:%s, expect:%s, result:%s\n", table.src, table.desc, res)
		}
	}
}
