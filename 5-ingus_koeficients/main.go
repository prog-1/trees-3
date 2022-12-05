package main

import "fmt"

func (n *node) len() int   { return n.val }
func (n *node) push(v int) { n.val++; insert(n, v) }
func (n *node) pop(v int)  { n.val--; delete(n, v) }

type node struct {
	val         int
	left, right *node
}

func insert(t *node, a int) *node {
	if t == nil {
		return &node{val: a}
	}
	if a < t.val {
		t.left = insert(t.left, a)

	} else {
		t.right = insert(t.right, a)
	}
	return t
}

func CreateBST(s []int) *node {
	var t *node
	for _, i := range s {
		t = insert(t, i)
	}
	return t
}

func find(t, parent *node, a int) (*node, *node) {
	if t == nil || t.val == a {
		return t, parent
	}
	if a < t.val {
		return find(t.left, t, a)
	}
	return find(t.right, t, a)
}

func next(t *node) *node {
	if t.left == nil && t.right == nil {
		return nil
	} else if t.left == nil {
		return t.right
	} else if t.right == nil {
		return t.left
	} else {
		tmp := t.right
		var parent *node
		for ; tmp.left != nil; tmp, parent = tmp.left, tmp {
		}
		t.val = tmp.val
		if t.val == t.right.val {
			t.right = nil
		} else {
			parent.left = nil
		}
		return t
	}
}

func delete(root *node, k int) *node {
	el, parent := find(root, root, k)
	if parent == el {
		next(el)
	} else if parent.left == el {
		parent.left = next(el)
	} else {
		parent.right = next(el)
	}
	return root
}

func inorder(t *node, ik int) (int, int) {
	if t == nil {
		return m, ik
	}
	inorder(t.left, f)
	if ik >= t.val {
		ik++
		m--
	}
	inorder(t.right, f)
}

func IngusKoef(m, ik int, ex []int) (int, int) {
	// t := CreateBST(ex[:m])
	var n *node
	for _, i := range ex {
		if m > 0 {
			n.len()
		}
	}
	return 0, 0
}

func main() {
	// n := 6
	m := 3
	ex := []int{3, 1, 7, 2, 6, 3} // len n
	//printTree(t)
	fmt.Println(IngusKoef(m, 1, ex))
}

func printTree(n *node) {
	printer(n, nil)
}

func printNonemptyHistory(prev []int, printLast bool) {
	if len(prev) == 0 {
		panic("must not happen")
	}
	last := len(prev)
	if !printLast {
		last--
	}
	for _, p := range prev[:last] {
		if p > 0 {
			fmt.Print("| ")
		} else {
			fmt.Print("  ")
		}
	}
}

func directChildren(n *node) (num int) {
	if n == nil {
		return 0
	}
	if n.left != nil {
		num++
	}
	if n.right != nil {
		num++
	}
	return num
}

func printer(n *node, prev []int) {
	if n == nil {
		return
	}
	if len(prev) > 0 {
		printNonemptyHistory(prev, true /*printLast*/)
		fmt.Println()
		printNonemptyHistory(prev, false /*printLast*/)
		prev[len(prev)-1]--
	}
	fmt.Println("+->", n.val)
	prev = append(prev, directChildren(n))
	if n.left != nil {
		printer(n.left, prev)
	}
	if n.right != nil {
		printer(n.right, prev)
	}
}
