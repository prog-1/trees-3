package main

type node struct {
	val         int
	left, right *node
}

func duplicates(root *node) bool {
	num := make(map[int]int)
	var check func(root *node, num map[int]int) bool
	check = func(root *node, num map[int]int) bool {
		if root == nil {
			return false
		}
		num[root.val]++
		if num[root.val] > 1 {
			return true
		}
		var leftCheck, rightCheck bool
		if root.left != nil {
			leftCheck = check(root.left, num)
		}
		if root.right != nil {
			rightCheck = check(root.right, num)
		}
		return leftCheck || rightCheck
	}
	return check(root, num)
}
func main() {
}
