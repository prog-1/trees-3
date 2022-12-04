package main

import (
	"reflect"
	"testing"
)

func TestDublicates(t *testing.T) {
	type testCase struct {
		name    string
		getTree func() (root *bt)
		want    bool
	}
	testCases := []testCase{
		{"nil", func() *bt { return nil }, false},
		{"has", func() *bt {
			return nbt(10).
				l(nbt(10))
		}, true},
		{"hasn't", func() *bt {
			return nbt(10).
				l(nbt(5))
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
		getTree func() (root *bt)
		want    []int
	}
	testCases := []testCase{
		{"nil", func() *bt { return nil }, []int{}},
		{"one level", func() *bt {
			return nbt(10)
		}, []int{10}},
		{"multiple levels", func() *bt {
			return nbt(10).
				l(nbt(3).
					l(nbt(8)).
					r(nbt(1))).
				r(nbt(6))
		}, []int{10, 6, 8}},
		{"dublicates", func() *bt {
			return nbt(5).
				l(nbt(10)).
				r(nbt(10))
		}, []int{5, 10}},
	}
	for _, tc := range testCases {
		root := tc.getTree()
		if got := levelMax(root); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}

func TestIngusCoefficient(t *testing.T) {
	type Input struct {
		complexities []int
		mit          int
	}
	type Want struct {
		miic, days int
	}
	for _, tc := range []struct {
		name  string
		input Input
		want  Want
	}{
		{"1", Input{[]int{3, 1, 7, 2, 6, 3}, 3}, Want{2, 3}},
		{"2", Input{[]int{7, 6, 5, 4, 3}, 9}, Want{3, 1}},
	} {
		gotMiic, gotDays := IngusCoefficient(tc.input.mit, tc.input.complexities)
		if gotMiic != tc.want.miic {
			t.Errorf("Incorrect minimal initial Ingus coefficient: got = %v, want = %v", gotMiic, tc.want.miic)
		}
		if gotDays != tc.want.days {
			t.Errorf("Incorrect day count: got = %v, want = %v", gotDays, tc.want.days)
		}
	}
}
