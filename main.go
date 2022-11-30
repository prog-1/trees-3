package main

type node struct {
	val         int
	left, right *node
}
type flag struct{}

func duplicates(t *node, flags map[int]flag) bool {
	if t == nil {
		return false
	}
	_, k := flags[t.val]
	if k {
		return true
	}
	flags[t.val] = flag{}
	return duplicates(t.left, flags) || duplicates(t.right, flags)
}

func Duplicates(t *node) bool {
	return duplicates(t, make(map[int]flag))
}

func inorderTraversal(root *node, elements map[int][]int, level int) int {
	if root == nil {
		return level - 1
	}
	elements[level] = append(elements[level], root.val)
	a := inorderTraversal(root.left, elements, level+1)
	b := inorderTraversal(root.right, elements, level+1)
	if a > b {
		return a
	}
	return b
}

func levelMax(t *node) []int {
	m := make(map[int][]int)
	l := inorderTraversal(t, m, 0)
	res := make([]int, l+1)
	for level, num := range m {
		res[level] = num[0]
		for _, n := range num[1:] {
			if res[level] < n {
				res[level] = n
			}
		}
	}
	return res
}

func main() {
}
