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
		l(n(2)).
		r(n(5))
	fmt.Println(hasDublicates(root))
}

func contains(s []int, val int) bool {
	for _, el := range s {
		if el == val {
			return true
		}
	}
	return false
}
func hasDublicates(root *node) bool {
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
