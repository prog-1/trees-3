package main

import "fmt"

type node struct {
	val         int
	left, right *node
}

func MaxLevelElements(t *node) []int {
	return FindMax(t, []int{}, 0)
}

func FindMax(t *node, max []int, level int) []int { // using preorder traversal
	if t == nil {
		return max
	}
	if len(max) < level+1 {
		max = append(max, t.val)
	} else if max[level] < t.val {
		max[level] = t.val
	}
	max = FindMax(t.left, max, level+1)
	return FindMax(t.right, max, level+1)
}

func main() {
	t := &node{val: 5, left: &node{val: 2, left: &node{val: 1}, right: &node{val: 3}}, right: &node{val: 7, right: &node{val: 9, left: &node{val: 8}}}}
	fmt.Println(MaxLevelElements(t))
}
