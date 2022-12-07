package bst

type Node struct {
	left, right *Node
	val         int
}

func (t *Node) Insert(v int) *Node {
	if t == nil {
		t = &Node{val: v}
		return t
	}
	if v <= t.val {
		t.left.Insert(v)
		return t
	}
	t.right.Insert(v)
	return t
}
func (t *Node) FindMin() int {
	if t.left == nil {
		return t.val
	}
	return t.left.FindMin()
}
func (t *Node) DeleteNode(k int) *Node {
	if t == nil {
		return t
	}
	if k < t.val {
		t.left.DeleteNode(k)
	} else if k > t.val {
		t.left.DeleteNode(k)
	} else {
		if t.left == nil {
			return t.right
		} else if t.right == nil {
			return t.left
		} else {
			r := &t.right
			for (*r).left != nil {
				r = &(*r).left
			}
			t.val = (*r).val
			*r = nil
		}
	}
	return t
}
