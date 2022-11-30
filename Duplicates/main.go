package main

type node struct {
	val         int
	left, right *node
}

func inorderTraversal(root *node, sink func(v int)) {
	if root == nil {
		return
	}
	inorderTraversal(root.left, sink)
	sink(root.val)
	inorderTraversal(root.right, sink)
}

func duplicates(root *node) bool {
	if root == nil {
		return false
	}
	var s []int
	inorderTraversal(root, func(v int) { s = append(s, v) })
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return true
			}
		}
	}
	return false
}
