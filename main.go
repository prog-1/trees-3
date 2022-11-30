package main

import "fmt"

type node struct {
	val         int
	left, right *node
}

func n(v int) *node                { return &node{v, nil, nil} }
func (root *node) l(n *node) *node { root.left = n; return root }
func (root *node) r(n *node) *node { root.right = n; return root }

func main() {
	root := n(10).
		l(n(2).
			l(n(1)).
			r(n(9))).
		r(n(5).
			l(n(1)).
			r(n(8)))

	fmt.Println(levelMax(root))
}

func hasDublicates(root *node) bool {
	contains := func(s []int, val int) bool {
		for _, el := range s {
			if el == val {
				return true
			}
		}
		return false
	}
	var values []int
	var hd func(*node) bool
	hd = func(r *node) bool {
		if r == nil {
			return false
		}
		if contains(values, r.val) {
			return true
		}
		values = append(values, r.val)
		return hd(r.left) || hd(r.right)
	}
	return hd(root)
}

func levelMax(root *node) []int {
	var values [][]int
	var traverse func(*node, int)
	traverse = func(root *node, lvl int) {
		if root == nil {
			return
		}
		if len(values) <= lvl {
			tmp := make([]int, 1)
			values = append(values, tmp)
		}
		values[lvl] = append(values[lvl], root.val)
		traverse(root.left, lvl+1)
		traverse(root.right, lvl+1)
	}
	traverse(root, 0)

	lm := make([]int, len(values))
	for i, lvl := range values {
		for _, val := range lvl {
			if lm[i] < val {
				lm[i] = val
			}
		}
	}

	return lm
}
