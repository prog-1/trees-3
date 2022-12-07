package main

type node struct {
	val         int
	left, right *node
}

func maxlevel(root *node) (res []int) {
	s := make([]int, 0)
	var find func(root *node, i int, s []int) []int
	find = func(root *node, i int, s []int) []int {
		if root == nil {
			return s
		}
		if i+1 > len(s) {
			s = append(s, root.val)
		}
		if root.val > s[i] {
			s[i] = root.val
		}
		if root.left != nil {
			s = find(root.left, i+1, s)
		}
		if root.right != nil {
			s = find(root.right, i+1, s)
		}
		return s
	}
	return find(root, 0, s)

}
func main() {
}
