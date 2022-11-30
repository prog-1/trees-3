package main

type node struct {
	val         int
	left, right *node
}

func duplicates(root *node, s []int) bool {
	if root == nil {
		return false
	}
	for i := range s {
		if root.val == s[i] {
			return true
		}
	}
	s = append(s, root.val)
	return duplicates(root.left, s) || duplicates(root.right, s)

}

// func main() {
// 	var s []int
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
// 		}}, s))
// }
