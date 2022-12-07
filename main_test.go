package main

import (
	"testing"
)

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
		//{"1", Input{[]int{3, 1, 7, 2, 6, 3}, 3}, Want{2, 3}},
		//{"2", Input{[]int{7, 6, 5, 4, 3}, 9}, Want{3, 1}},
		{"3", Input{[]int{7, 6, 2, 1, 1 /*, 17, 69, 72*/}, 4}, Want{3, 1}},
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
