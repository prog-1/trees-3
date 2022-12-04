package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func treeToSlice[T any](root *node[T]) (values []T, counts []int) {
	var fillSlice func(*node[T])
	fillSlice = func(cur *node[T]) {
		if cur == nil {
			return
		}
		values = append(values, cur.val)
		counts = append(counts, cur.count)

		fillSlice(cur.left)
		fillSlice(cur.right)
	}
	fillSlice(root)
	return values, counts
}

func TestDublicates(t *testing.T) {
	type testCase struct {
		name    string
		getTree func() (root *node[int])
		want    bool
	}
	testCases := []testCase{
		{"nil", func() *node[int] { return nil }, false},
		{"has", func() *node[int] {
			return n(10, 1).
				l(n(10, 1))
		}, true},
		{"hasn't", func() *node[int] {
			return n(10, 1).
				l(n(5, 1))
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
		getTree func() (root *node[int])
		want    []int
	}
	testCases := []testCase{
		{"nil", func() *node[int] { return nil }, []int{}},
		{"one level", func() *node[int] {
			return n(10, 1)
		}, []int{10}},
		{"multiple levels", func() *node[int] {
			return n(10, 1).
				l(n(3, 1).
					l(n(8, 1)).
					r(n(1, 1))).
				r(n(6, 1))
		}, []int{10, 6, 8}},
		{"dublicates", func() *node[int] {
			return n(5, 1).
				l(n(10, 1)).
				r(n(10, 1))
		}, []int{5, 10}},
	}
	for _, tc := range testCases {
		root := tc.getTree()
		if got := levelMax(root); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}

func TestWordCount(t *testing.T) {
	for _, tc := range []struct {
		name     string
		filename string
		wantTree func() *node[string]
	}{
		{"1", "text.txt", func() *node[string] {
			return n("Only", 1).
				l(n("Albert", 1).
					r(n("Einstein", 1))).
				r(n("a", 2).
					r(n("life", 2).
						l(n("for", 1).
							r(n("is", 1))).
						r(n("lived", 1).
							r(n("others", 1).
								r(n("worthwhile", 1))))))
		}},
	} {
		file, err := ioutil.ReadFile(tc.filename)
		if err != nil {
			panic(err)
		}
		gotValues, gotCounts := treeToSlice(countWords(file))
		wantValues, wantCounts := treeToSlice(tc.wantTree())
		if !reflect.DeepEqual(gotValues, wantValues) {
			t.Errorf("Incorrect values: got = %v, want = %v", gotValues, wantValues)
		}
		if !reflect.DeepEqual(gotCounts, wantCounts) {
			t.Errorf("Incorrect counts: got = %v, want = %v", gotCounts, wantCounts)
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
