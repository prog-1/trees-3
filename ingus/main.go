package main

import "fmt"

type node struct {
	val         int
	left, right *node
}

func minVal(root *node) int {
	for root.left != nil {
		root = root.left
	}
	return root.val
}

func deleteNode(root *node, k int) *node {
	if root == nil {
		return root
	}
	if root.val > k {
		root.left = deleteNode(root.left, k)
		return root
	} else if root.val < k {
		root.right = deleteNode(root.right, k)
		return root
	}
	if root.left == nil {
		return root.right
	} else if root.right == nil {
		return root.left
	}
	parent := root
	res := root.right
	for res.left != nil {
		parent = res
		res = res.left
	}
	if parent != root {
		parent.left = res.right
	} else {
		parent.right = res.right
	}
	root.val = res.val
	return root
}

func findIK(n, m int, s []int) (int, int) {
	tree := &node{val: s[0], left: nil, right: nil}
	for i := 1; i < m; i++ {
		insert(tree, s[i])
	}
	s = s[m:]
	res, ik := 1, 1
	days := 0
	for len(s) > 0 || tree != nil {
		min := minVal(tree)
		if ik >= min {
			ik += min
			tree = deleteNode(tree, min)
			if len(s) > 0 {
				insert(tree, s[0])
				s = s[1:]
				days++
			}
		} else {
			res += (min - ik)
		}
	}
	return res, days
}

func insert(root *node, key int) {
	nodes := &node{val: key, left: nil, right: nil}
	if root == nil {
		return
	}
	prev := &node{val: key, left: nil, right: nil}
	tmp := root
	for tmp != nil {
		if tmp.val > key {
			prev = tmp
			tmp = tmp.left
		} else if tmp.val < key {
			prev = tmp
			tmp = tmp.right
		}
	}
	if prev.val > key {
		prev.left = nodes
	} else {
		prev.right = nodes
	}
}

// func inorder(t *node, f func(int)) {
// 	if t == nil {
// 		return
// 	}
// 	inorder(t.left, f)
// 	f(t.val)
// 	inorder(t.right, f)
// }

func main() {
	// tree := &node{left: nil, right: nil}
	// insert(1, tree)
	// inorder(tree, func(v int) { fmt.Println(v) })
	min, days := findIK(6, 3, []int{3, 1, 7, 2, 6, 3})
	fmt.Println(min, days)
}
