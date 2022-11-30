package main

import (
	"reflect"
	"testing"
)

func TestLevelMax(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    *node
		want []int
	}{
		{"test1", t1(), []int{}},
		{"test2", t2(), []int{7, 9, 10}},
		{"test3", t3(), []int{9, 9, 10}},
		{"test4", t4(), []int{7, 9, 10, 11}},
		{"test5", t5(), []int{7, 9, 10, 32}},
		{"test6", t6(), []int{1}},
	} {

		t.Run(tc.name, func(t *testing.T) {
			if got := levelMax(tc.n); !reflect.DeepEqual(got, tc.want) {
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
				right: &node{val: 4}}},

		right: &node{
			val: 9,
			left: &node{val: 8,
				left:  &node{val: 11},
				right: &node{val: 32}},
			right: &node{val: 10},
		}}
}
func t6() *node {
	return &node{val: 1,
		left:  nil,
		right: nil}
}
