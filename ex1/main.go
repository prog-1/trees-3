package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	cat       int
	neighbors []*node
}

//Doesnt work with trees that haveonly reversed pathes like test 8 on codeforce

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
		if b < a {
			continue
		}
		nodes[a-1].neighbors = append(nodes[a-1].neighbors, &nodes[b-1])

	}
	if len(nodes) < n {
		fmt.Println(0)
	}
	return &nodes[0], m
}
func roots(root *node, m int) int {
	var visit func(root *node, catcnt int, ways, m int, cnt *int) int
	visit = func(root *node, catcnt int, ways, m int, cnt *int) int {
		if root == nil {
			return 0
		}

		if root.cat == 0 {
			catcnt = 0
		} else {
			catcnt += root.cat
			if catcnt > m {
				return ways
			}
		}
		if len(root.neighbors) == 0 && *cnt > 0 {
			return ways + 1
		}
		*cnt++
		for _, v := range root.neighbors {
			ways = visit(v, catcnt, ways, m, cnt)
		}
		return ways
	}
	var cnt int
	return visit(root, 0, 0, m, &cnt)

}
func main() {
	n, m := collectData()
	fmt.Println(roots(n, m))
}
