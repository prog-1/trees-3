package main

import (
	"testing"
)

func TestDuplicates(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    *node
		want bool
	}{
		{"test1", t1(), false},
		{"test2", t2(), false},
		{"test3", t3(), true},
		{"test4", t4(), false},
		{"test5", t5(), true},
		{"test6", t6(), false},
	} {

		t.Run(tc.name, func(t *testing.T) {
			if got := duplicates(tc.n); !got == tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func t1() *node {
	return nil
}
func t2() *node {
	return &node{val: 7,
		left: &node{
			val:   5,
			left:  &node{val: 1},
			right: &node{val: 6},
		},
		right: &node{
			val:   9,
			left:  &node{val: 8},
			right: &node{val: 10},
		}}
}
func t3() *node {
	return &node{val: 9,
		left: &node{
			val:   5,
			left:  &node{val: 1},
			right: &node{val: 6},
		},
		right: &node{
			val:   9,
			left:  &node{val: 8},
			right: &node{val: 10},
		}}
}
func t4() *node {
	return &node{val: 7,
		left: &node{
			val:   5,
			left:  &node{val: 1},
			right: &node{val: 6},
		},
		right: &node{
			val: 9,
			left: &node{val: 8,
				left:  &node{val: 11},
				right: &node{val: 4}},
			right: &node{val: 10},
		}}
}
func t5() *node {
	return &node{val: 7,
		left: &node{
			val:  5,
			left: &node{val: 1},
			right: &node{val: 6,
				left:  &node{val: 11},
				right: &node{val: 5}}},

		right: &node{
			val: 9,
			left: &node{val: 8,
				left:  &node{val: 11},
				right: &node{val: 4}},
			right: &node{val: 10},
		}}
}
func t6() *node {
	return &node{val: 0,
		left:  nil,
		right: nil}
}
