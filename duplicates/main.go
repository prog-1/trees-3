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

//second task: levelmax

//inpired by anton's solution
func levelMax(t *node, level int, s []int) []int {
	if t == nil {
		return s
	}
	if level == len(s) {
		s = append(s, t.val)
	}
	s[level] = findMax(t.val, s[level])
	if t.left != nil {
		s = levelMax(t.left, level+1, s)
	}
	if t.right != nil {
		s = levelMax(t.right, level+1, s)
	}
	return s
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	m := make(map[int]bool)
	tree := &node{val: 4, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 6, left: &node{val: 5}, right: &node{val: 7}}}
	fmt.Println(duplicates(tree, m))
	fmt.Println(levelMax(tree, 0, nil))
}
