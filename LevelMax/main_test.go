package main

import (
	"reflect"
	"testing"
)

func TestLevelMax(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *node
		want []int
	}{
		{"case-1", nil, nil},
		{"case-2", &node{val: 1}, []int{1}},
		{"case-3", &node{val: 5, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 7}}, []int{5, 7, 3}},
		{"case-4", &node{val: 5, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 7, right: &node{val: 9, left: &node{val: 8}}}}, []int{5, 7, 9, 8}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := levelMax(tc.root); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
