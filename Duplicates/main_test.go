package main

import "testing"

func TestDuplicates(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *node
		want bool
	}{
		{"case-1", nil, false},
		{"case-2", &node{val: 1}, false},
		{"case-3", &node{val: 1, left: &node{val: 2}, right: &node{val: 3}}, false},
		{"case-4", &node{val: 1, left: &node{val: 2}, right: &node{val: 1}}, true},
		{"case-5", &node{val: 1, left: &node{val: 2, left: &node{val: 1}}}, true},
		{"case-6", &node{val: 1, left: &node{val: 1, left: &node{val: 1}}}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := duplicates(tc.root); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
