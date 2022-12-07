package main

import "testing"

func TestDuplicates(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *node
		want bool
	}{
		{"case-1", nil, false},
		{"case-2", &node{}, false},
		{"case-3", &node{val: 4}, false},
		{"case-4", &node{val: 4, left: &node{val: 9}}, false},
		{"case-5", &node{val: 4, right: &node{val: 9}}, false},
		{"case-6", &node{val: 5, right: &node{val: 3}}, false},
		{"case-7", &node{val: 4, left: &node{val: 4}}, true},
		{"case-8", &node{val: 4, right: &node{val: 4}}, true},
		{"case-9", &node{val: 6, right: &node{val: 6}}, true},
		{"case-10", &node{val: 7, left: &node{val: 5, left: &node{val: 1}}, right: &node{val: 9, left: &node{val: 8}, right: &node{val: 10}}}, false},
		{"case-11", &node{val: 7, left: &node{val: 5, left: &node{val: 1}}, right: &node{val: 9, left: &node{val: 1}, right: &node{val: 10}}}, true},
		{"case-12", &node{val: 7, left: &node{val: 5, left: &node{val: 1}}, right: &node{val: 9, left: &node{val: 5}, right: &node{val: 10}}}, true},
		{"case-13", &node{val: 7, left: &node{val: 5, left: &node{val: 1}}, right: &node{val: 9, left: &node{val: 8}, right: &node{val: 9}}}, true},
		{"case-14", &node{val: 16, left: &node{val: 16, left: &node{val: 16}}, right: &node{val: 15, right: &node{val: 16}}}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Duplicates(tc.root); got != tc.want {
				t.Errorf("Duplicates(%v) = %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}
