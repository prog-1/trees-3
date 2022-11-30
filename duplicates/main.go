package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func duplicates(t *node, m map[int]bool) bool {
	if t == nil {
		return false
	}
	if _, ok := m[t.val]; ok {
		return true
	}
	m[t.val] = true
	return duplicates(t.left, m) || duplicates(t.right, m)
}

func main() {
	m := make(map[int]bool)
	tree := &node{val: 4, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 6, left: &node{val: 5}, right: &node{val: 7}}}
	fmt.Println(duplicates(tree, m))
}
