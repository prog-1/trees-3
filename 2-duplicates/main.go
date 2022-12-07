package main

import "fmt"

type node struct {
	val         int
	left, right *node
}

func Duplicates(t *node) bool {
	return CheckDuplicates(t, map[int]int{})
}

func CheckDuplicates(t *node, m map[int]int) bool { // returns true if there are duplicates
	if t == nil {
		return false
	}
	if _, ok := m[t.val]; ok {
		return true
	} else {
		m[t.val]++
		return CheckDuplicates(t.left, m) || CheckDuplicates(t.right, m)
	}
}

func main() {
	t := &node{val: 7, left: &node{val: 5, left: &node{val: 1}}, right: &node{val: 9, left: &node{val: 8}, right: &node{val: 10}}}
	fmt.Println(Duplicates(t))
	t = &node{val: 7, left: &node{val: 5, left: &node{val: 1}}, right: &node{val: 9, left: &node{val: 1}, right: &node{val: 10}}}
	fmt.Println(Duplicates(t))
}
