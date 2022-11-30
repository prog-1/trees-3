package main

import (
	"bytes"
	"reflect"
	"testing"
)

func tree01() *node {
	return nil
}

func tree02() *node {
	return &node{val: 1}
}

func tree03() *node {
	return &node{
		val: 7,
		left: &node{
			val: 3,
		},
		right: &node{
			val: 8,
		},
	}
}
func tree04() *node {
	return &node{
		val: -4,
		left: &node{
			val: -4,
		},
		right: &node{
			val: -3,
		},
	}
}

func TestLevelMax(t *testing.T) {
	for _, tc := range []struct {
		name string
		root *node
		want []int
	}{
		{tree01(), []int{}},
		{tree02(), []int{1}},
		{tree03(), []int{7, 8}},
		{tree04(), []int{-4, -3}},
	} {

		if got := levelMax(tc.root); !reflect.DeepEqual(got, tc.want) {
			var buf bytes.Buffer
			t.Errorf("(#%v) = %v, want %v\n%v", tc.name, got, tc.want, buf.String())
		}
	}
}
