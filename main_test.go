package main

import "testing"
import "reflect"

func TestDublicates(t *testing.T) {
	type testCase struct {
		name    string
		getTree func() (root *node)
		want    bool
	}
	testCases := []testCase{
		{"nil", func() *node { return nil }, false},
		{"has", func() *node {
			return n(10).
				l(n(10))
		}, true},
		{"hasn't", func() *node {
			return n(10).
				l(n(5))
		}, false},
	}
	for _, tc := range testCases {
		root := tc.getTree()
		if got := hasDublicates(root); got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}

func TestLevelMax(t *testing.T) {
	type testCase struct {
		name    string
		getTree func() (root *node)
		want    []int
	}
	testCases := []testCase{
		{"nil", func() *node { return nil }, []int{}},
		{"one level", func() *node {
			return n(10)
		}, []int{10}},
		{"multiple levels", func() *node {
			return n(10).
				l(n(3).
					l(n(8)).
					r(n(1))).
				r(n(6))
		}, []int{10, 6, 8}},
		{"dublicates", func() *node {
			return n(5).
				l(n(10)).
				r(n(10))
		}, []int{5, 10}},
	}
	for _, tc := range testCases {
		root := tc.getTree()
		if got := levelMax(root); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
