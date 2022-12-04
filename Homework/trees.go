package main

type Node struct {
	val         int
	left, right *Node
}

func n(v int) *Node             { return &Node{val: v} }
func (n *Node) l(l *Node) *Node { n.left = l; return n }
func (n *Node) r(r *Node) *Node { n.right = r; return n }

//Trees

//[12,10,5]
func Tree1() *Node {
	return n(12).
		l(n(8).
			l(n(5)).
			r(n(1))).
		r(n(10).
			l(n(1)).
			r(n(4)))
}

func Tree2() *Node {
	return n(8).
		l(n(3).
			l(n(1).
				r(n(6).
					l(n(4)).
					r(n(7)))).
			r(n(10).
				l(n(9).
					l(n(2)).
					r(n(12))).
				r(n(14).
					l(n(13)))))
}

func Tree3() *Node {
	return n(5).l(n(6)).r(n(5))
}

func Tree4() *Node {
	return n(5).l(n(6))
}

func Tree5() *Node {
	return n(0)
}

func Tree6() *Node {
	return nil
}
