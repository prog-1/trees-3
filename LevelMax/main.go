package main

type node struct {
	val         int
	left, right *node
}

func levelMax(root *node) (res []int) {
	if root == nil {
		return nil
	}
	m := levelOrderTraversal(root)
	res = append(res, m[0][0].val)
	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[i]); j++ {
			m[i][j].val = max(m[i][j-1].val, m[i][j].val)
		}
		res = append(res, m[i][len(m[i])-1].val)
	}
	return
}

func levelOrderTraversal(root *node) map[int][]*node {
	if root == nil {
		return nil
	}
	m := make(map[int][]*node)
	m[0] = append(m[0], root)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j].left != nil {
				m[i+1] = append(m[i+1], m[i][j].left)
			}
			if m[i][j].right != nil {
				m[i+1] = append(m[i+1], m[i][j].right)
			}
		}
	}
	return m
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
