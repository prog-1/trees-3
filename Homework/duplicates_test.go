package main

import "testing"

func TestCheckDuplicates(t *testing.T) {
	for _, tc := range []struct {
		name string
		t    *Node
		want bool
	}{
		{name: "Has dublicates 1", t: Tree1(),
			want: true},
		{name: "No dublicates 1", t: Tree2(),
			want: false},
		{name: "Has dublicates 2", t: Tree3(),
			want: true},
		{name: "No dublicates 2", t: Tree4(),
			want: false},
		{name: "One element", t: Tree5(),
			want: false},
		{name: "nil", t: Tree6(),
			want: false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := CheckDuplicates(tc.t)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
