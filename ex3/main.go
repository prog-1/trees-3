package main

type node struct {
	val         int
	left, right *node
}

func levelmax(root *node) []int {
	s := make([]int, 0)
	var visit func(root *node, l int, s []int) []int
	visit = func(root *node, l int, s []int) []int {
		if root == nil {
			return s
		}
		if l+1 > len(s) {
			s = append(s, root.val)
		}
		if root.val > s[l] {
			s[l] = root.val
		}
		if root.left != nil {
			s = visit(root.left, l+1, s)
		}
		if root.right != nil {
			s = visit(root.right, l+1, s)
		}
		return s
	}
	return visit(root, 0, s)
}

// func main() {
// 	var s []int
// 	fmt.Println(levelmax(&node{val: 7,
// 		left: &node{
// 			val:   5,
// 			left:  &node{val: 1},
// 			right: &node{val: 6},
// 		},
// 		right: &node{
// 			val:   9,
// 			left:  &node{val: 8},
// 			right: &node{val: 10},
// 		}}, 0, s))
// }
