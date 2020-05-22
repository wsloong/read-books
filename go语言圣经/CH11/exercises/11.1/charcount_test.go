package charcount

import "testing"

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input string
	}{
		{""},
		{"123"},
		{"abc"},
		{"123abc"},
		{"1f2z3'"},
		{"*23()!@#"},
	}

	for _, test := range tests {
		t.Logf("*******input:%v*********\n", test.input)
		CharCount(test.input)
	}

}
