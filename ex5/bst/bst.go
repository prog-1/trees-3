package bst

type Node struct {
	Left, Right *Node
	Val         int
}

func (t *Node) NodeCount() int {
	if t == nil {
		return 0
	}
	return t.Left.NodeCount() + t.Right.NodeCount() + 1
}
func (t *Node) Insert(v int) *Node {
	if t == nil {
		t = &Node{Val: v}
		return t
	}
	if v <= t.Val {
		t.Left = t.Left.Insert(v)
		return t
	}
	t.Right = t.Right.Insert(v)
	return t
}
func (t *Node) FindMin() int {
	// fmt.Println("val=", t.Val)
	if t.Left == nil {
		return t.Val
	}
	return t.Left.FindMin()
}
func (t *Node) FindMax() int {
	// fmt.Println("val=", t.Val)
	if t.Right == nil {
		return t.Val
	}
	return t.Right.FindMax()
}
func (t *Node) DeleteNode(k int) *Node {
	if t == nil {
		return t
	}
	if k < t.Val {
		t.Left = t.Left.DeleteNode(k)
	} else if k > t.Val {
		t.Right = t.Right.DeleteNode(k)
	} else {
		if t.Left == nil {
			return t.Right
		} else if t.Right == nil {
			return t.Left
		} else {
			r := &t.Right
			for (*r).Left != nil {
				r = &(*r).Left
			}
			t.Val = (*r).Val
			*r = nil
		}
	}

	return t
}
