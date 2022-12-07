package main

import "fmt"

type node struct {
	val         int
	left, right *node
}

func levelMax(root *node) (res []int) {
	output := make([]int, 0)
	var visit func(root *node, level int, output []int) []int
	visit = func(root *node, level int, output []int) []int {
		if root == nil {
			return output
		}
		if level+1 > len(output) {
			output = append(output, root.val)
		} else if root.val > output[level] {
			output[level] = root.val
		}
		if root.left != nil {
			output = visit(root.left, level+1, output)
		}
		if root.right != nil {
			output = visit(root.right, level+1, output)
		}
		return output
	}
	return visit(root, 0, output)

}
func main() {
	fmt.Println(levelMax(&node{val: 7,
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
