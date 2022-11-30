package main

import (
	"bytes"
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
			val: 7,
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

func TestDuplicates(t *testing.T) {
	for _, tc := range []struct {
		root *node
		want bool
	}{
		{tree01(), false},
		{tree02(), false},
		{tree03(), true},
		{tree04(), true},
	} {

		if got := duplicates(tc.root); got != tc.want {
			var buf bytes.Buffer
			t.Errorf("(#%v) = %v, want %v\n%v", tc.name, got, tc.want, buf.String())
		}
	}
}
