package main

import (
	"reflect"
	"testing"
)

func TestDuplicates(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *node
		want []int
	}{
		{"case-1", nil, []int{}},
		{"case-2", &node{val: 4}, []int{4}},
		{"case-3", &node{val: 4, left: &node{val: 9}}, []int{4, 9}},
		{"case-4", &node{val: 4, right: &node{val: 9}}, []int{4, 9}},
		{"case-5", &node{val: 5, right: &node{val: 3}}, []int{5, 3}},
		{"case-6", &node{val: 4, left: &node{val: 4}}, []int{4, 4}},
		{"case-7", &node{val: 4, right: &node{val: 4}}, []int{4, 4}},
		{"case-8", &node{val: 6, right: &node{val: 6}}, []int{6, 6}},
		{"case-9", &node{val: 7, left: &node{val: 5, left: &node{val: 1}}, right: &node{val: 9, left: &node{val: 8}, right: &node{val: 10}}}, []int{7, 9, 10}},
		{"case-10", &node{val: 7, left: &node{val: 5, left: &node{val: 1}}, right: &node{val: 9, left: &node{val: 9}, right: &node{val: 7}}}, []int{7, 9, 9}},
		{"case-11", &node{val: 7, left: &node{val: 5, left: &node{val: 1}}, right: &node{val: 9, left: &node{val: 5}, right: &node{val: 5}}}, []int{7, 9, 5}},
		{"case-12", &node{val: 7, left: &node{val: 11, left: &node{val: 1}}, right: &node{val: 9, left: &node{val: 10}, right: &node{val: 7}}}, []int{7, 11, 10}},
		{"case-13", &node{val: 16, left: &node{val: 16, left: &node{val: 16}}, right: &node{val: 15, right: &node{val: 16}}}, []int{16, 16, 16}},
		{"example-1", &node{val: 5, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 7, right: &node{val: 9, left: &node{val: 8}}}}, []int{5, 7, 9, 8}},
		{"example-2", &node{val: 5, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 7}}, []int{5, 7, 3}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaxLevelElements(tc.root); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Duplicates(%v) = %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}
