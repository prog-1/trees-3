package main

import (
	"reflect"
	"testing"
)

func TestLevelMax(t *testing.T) {
	for _, tc := range []struct {
		name string
		t    *Node
		want []int
	}{
		{name: "Tree 1", t: Tree1(), want: []int{12, 10, 5}},
		{name: "Tree 2", t: Tree2(), want: []int{8, 3, 10, 14, 13}},
		{name: "Tree 3", t: Tree3(), want: []int{5, 6}},
		{name: "Tree 4", t: Tree4(), want: []int{5, 6}},
		{name: "Tree with one element", t: Tree5(), want: []int{0}},
		{name: "Nil Tree", t: Tree6(), want: nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := LevelMax(tc.t)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
