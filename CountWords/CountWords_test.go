package main

import (
	"testing"
)

func TestCreateBSTFromString(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  map[string]int
	}{
		{"a b", map[string]int{"a": 1, "b": 1}},
		{"a, b", map[string]int{"a": 1, "b": 1}},
		{"a, b a", map[string]int{"a": 2, "b": 1}},
		{"", map[string]int{}},
		{"a!!!!a", map[string]int{"a": 2}},
	} {

		got := createBSTFromString([]byte(tc.input))
		if !Contains_exactly(got, tc.want) {
			t.Errorf("createBSTFromString(\"%v\")=%v, want = %v", tc.input, got, tc.want)
		}
	}
}

func Contains_exactly(t *node, m map[string]int) bool {
	inorderTraversal(t, func(v *node) {
		m[string(v.word)] -= v.repeates
	})
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
