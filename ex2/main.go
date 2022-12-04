package main

type node struct {
	val         int
	left, right *node
}

func duplicates(root *node) bool {
	s := make([]int, 0)
	var visit func(root *node, s []int) bool
	visit = func(root *node, s []int) bool {
		if root == nil {
			return false
		}
		for i := range s {
			if root.val == s[i] {
				return true
			}
		}
		s = append(s, root.val)
		return visit(root.left, s) || visit(root.right, s)
	}

	return visit(root, s)

}

// func main() {
// 	fmt.Println(duplicates(&node{val: 9,
// 		left: &node{
// 			val:   5,
// 			left:  &node{val: 1},
// 			right: &node{val: 6},
// 		},
// 		right: &node{
// 			val:   9,
// 			left:  &node{val: 8},
// 			right: &node{val: 10},
// 		}}))
// }
