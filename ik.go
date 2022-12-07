package main

import (
	"fmt"
	"reflect"
)

type node struct {
	val         int
	left, right *node
}

// "solve" function not working
func solve(m int, d []int, minIK int) (int, int) {
	var (
		s    []int
		days int
	)
	t, ik, d, lenT := &node{val: d[0]}, minIK, d[1:], 1
	for len(d) != 0 {
		t = bst(m-lenT, d, t)
		lenT = m
		fmt.Println(t, t.left, t.right, ik)
		inorderTraversal(t, func(v int) { s = append(s, v) })
		tmp := s
		for i := 0; i < len(tmp); i++ {
			if ik >= tmp[i] {
				ik++
				delete(t, tmp[i])
				tmp, i, lenT = append(tmp[:i], tmp[i+1:]...), i-1, lenT-1
			}
		}
		if !reflect.DeepEqual(tmp, s) {
			minIK++
			minIK, days = solve(m, d, minIK)
		}
		days++
		d = d[m+1:]
	}
	return minIK, days
}

func bst(m int, d []int, t *node) *node {
	for i := 0; i < m-1; i++ {
		if d[i] < t.val {
			if t.left != nil {
				bst((m - i), d, t.left)
			} else {
				t.left = new(node)
				t.left.val = d[i]
			}
		} else {
			if t.right != nil {
				bst((m - i), d, t.right)
			} else {
				t.right = new(node)
				t.right.val = d[i]
			}
		}
	}
	return t
}

func inorderTraversal(root *node, sink func(v int)) {
	if root == nil {
		return
	}
	inorderTraversal(root.left, sink)
	sink(root.val)
	inorderTraversal(root.right, sink)
}

func delete(root *node, k int) *node {
	if root == nil {
		return nil
	}
	if k < root.val {
		root.left = delete(root.left, k)
	} else if k > root.val {
		root.right = delete(root.right, k)
	} else {
		if root.left == nil && root.right == nil {
			return nil
		} else if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			tmp := root.right
			for tmp.left != nil {
				tmp = tmp.left
			}
			root.val = tmp.val
			root.right = delete(root.right, root.val)
		}
	}
	return root
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&d[i])
	}
	fmt.Println(solve(m, d, 1))
}
