package main

import (
	"reflect"
	"testing"
)

func small() *node {
	return &node{val: 3, left: &node{val: 2, left: &node{val: 1}}, right: &node{val: 4}}
}

func large() *node {
	return &node{val: 4, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 6, left: &node{val: 5}, right: &node{val: 7}}}
}

func TestDuplicates(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input *node
		want  bool
	}{
		{"nil", nil, false},
		{"small", &node{val: 3, left: &node{val: 2, left: &node{val: 1}}, right: &node{val: 4}}, false},
		{"small, duplicates", &node{val: 3, left: &node{val: 2, left: &node{val: 4}}, right: &node{val: 4}}, true},
		{"small, same val", &node{val: 1, left: &node{val: 1, left: &node{val: 1}}, right: &node{val: 1}}, true},
		{"large", large(), false},
		{"large, duplicates", &node{val: 4, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 6, left: &node{val: 5}, right: &node{val: 1}}}, true},
		{"large, duplicates 2", &node{val: 4, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 4}}, right: &node{val: 6, left: &node{val: 5}, right: &node{val: 7}}}, true},
		{"large, duplicates 3", &node{val: 4, left: &node{val: 7, left: &node{val: 1}, right: &node{val: 4}}, right: &node{val: 6, left: &node{val: 5}, right: &node{val: 7}}}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := make(map[int]bool)
			if got := duplicates(tc.input, m); got != tc.want {
				t.Errorf("Test: %v, got = %v, want = %v", tc.name, got, tc.want)
			}
		})
	}
}

func TestLevelMax(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input *node
		want  []int
	}{
		{"nil", nil, nil},
		{"small", small(), []int{3, 4, 1}},
		{"large", large(), []int{4, 6, 7}},
		{"small, duplicates", &node{val: 4, left: &node{val: 2, left: &node{val: 4}}, right: &node{val: 4}}, []int{4, 4, 4}},
		{"large duplicates", &node{val: 1, left: &node{val: 3, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 3, left: &node{val: 3}, right: &node{val: 7}}}, []int{1, 3, 7}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := levelMax(tc.input, 0, nil); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Test: %v, got = %v, want = %v", tc.name, got, tc.want)
			}
		})
	}
}
