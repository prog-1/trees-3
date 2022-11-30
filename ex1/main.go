package main

//https://codeforces.com/contest/580/submission/183405832

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	cat       int
	neighbors []*node
	// prev      *node
}

func collectData() (*node, int) {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n, &m)
	nodes := make([]node, n)
	for i := range nodes {
		fmt.Fscan(r, &nodes[i].cat)

	}
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(r, &a, &b)
		nodes[a-1].neighbors = append(nodes[a-1].neighbors, &nodes[b-1])
		nodes[b-1].neighbors = append(nodes[b-1].neighbors, &nodes[a-1])
		// if a > b {
		// 	nodes[a-1].prev = &nodes[b-1]
		// } else {
		// 	nodes[b-1].prev = &nodes[a-1]
		// }
	}
	return &nodes[0], m
}
func roots(root *node, m int) int {
	var visit func(root, prev *node, catcnt int, ways, m int) int
	visit = func(root, prev *node, catcnt int, ways, m int) int {
		if root.cat == 0 {
			catcnt = 0
		} else {
			catcnt += root.cat
			if catcnt > m {
				return ways
			}
		}
		if len(root.neighbors) == 1 && prev != nil {
			return ways + 1
		}
		for _, v := range root.neighbors {
			if v != prev {
				ways = visit(v, root, catcnt, ways, m)
			}
		}
		return ways
	}
	return visit(root, nil, 0, 0, m)

}
func main() {
	n, m := collectData()
	fmt.Println(roots(n, m))

}
