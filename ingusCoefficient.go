package main

import "fmt"

type node struct {
	val, count  int
	left, right *node
}

func main() {
	dta := 9
	book := []int{28, 59, 87, 19, 19, 17, 69, 72}
	fmt.Println(ingusCoefficient(book, dta))
}

func ingusCoefficient(book []int, dta int) int {
	var ic, mic, nt, st int
	var t *node
	var min *node
	if dta > len(book) {
		dta = len(book)
	}

	for len(book) != 0 || t.len() != 0 {
		if min != nil && ic >= min.val {
			delMin(&t)
			ic++
			st++
			min = t.min()
		} else if t.len() < dta && len(book) != 0 {
			nt = dta - t.len()
			for i := 0; i < nt; i++ {
				t = t.add(book[i])
			}
			book = book[nt:]
			min = t.min()
		} else /*dt.len() == dta*/ {
			mic = min.val - st
			ic = mic + st
		}
	}
	return mic

}

/*
dta - day task amount
t - day task tree
nt - new tasks
st - solved tasks
*/

func (t *node) add(val int) *node {
	n := &t
	for *n != nil {
		if val < (*n).val {
			n = &(*n).left
		} else {
			n = &(*n).right
		}
	}
	*n = &node{val: val}
	return t
}

func delMin(t **node) {
	for (*t).left != nil {
		t = &(*t).left
	}
	*t = (*t).right
}

func (t *node) min() *node {
	if t == nil {
		return t
	}
	for ; t.left != nil; t = t.left {
	}
	return t
}

func (t *node) len() int {
	var l int
	if t == nil {
		return l
	} else if t.count == 0 {
		l++
	} else {
		for i := 0; i < t.count; i++ {
			l++
		}
	}
	return l + t.left.len() + t.right.len()
}
