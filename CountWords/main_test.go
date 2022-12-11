package main

import "testing"

func TestCountWords(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []string
		want *node
	}{
		{"case-1", nil, nil},
		{"case-2", []string{"hello"}, &node{val: "hello", count: 1}},
		{"case-3", []string{"hello", "world"}, &node{val: "hello", count: 1, right: &node{val: "world", count: 1}}},
		{"case-4", []string{"hello", "world", "hello"}, &node{val: "hello", right: &node{val: "world", count: 1}}},
		{"case-5", []string{"hello", "hello", "hello", "world", "world"}, &node{val: "hello", count: 3, right: &node{val: "world", count: 2}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := countWords(nil, tc.s); !equal(got, tc.want) {
				t.Errorf("got = %v, want =  %v", got, tc.want)
			}
		})
	}
}

func equal(a, b *node) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else if a.val != b.val {
		return false
	}
	return equal(a.left, b.left) && equal(a.right, b.right)
}
