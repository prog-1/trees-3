package main

import "testing"

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
