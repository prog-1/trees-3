package main

import "fmt"

type node struct {
	val         int
	left, right *node
}

func duplicates(root *node) bool {
	numbers := make(map[int]int)
	var visit func(root *node, numbers map[int]int) bool
	visit = func(root *node, numbers map[int]int) bool {
		if root == nil {
			return false
		}
		numbers[root.val]++
		if numbers[root.val] > 1 {
			return true
		}
		var a, b bool
		if root.left != nil {
			a = visit(root.left, numbers)
		}
		if root.right != nil {
			b = visit(root.right, numbers)
		}
		return a || b
	}

	return visit(root, numbers)

}
func main() {
	fmt.Println(duplicates(&node{val: 9,
		left: &node{
			val:   5,
			left:  &node{val: 1},
			right: &node{val: 6},
		},
		right: &node{
			val:   9,
			left:  &node{val: 8},
			right: &node{val: 10},
		}}))
}
