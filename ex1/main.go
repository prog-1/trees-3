package main

//https://codeforces.com/contest/580/submission/183837960

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type node struct {
	cat       int
	neighbors []*node
}

func data(r io.Reader) (*node, int) {
	var n, m int
	fmt.Fscan(r, &n, &m)
	nodes := make([]node, n)
	for i := range nodes {
		fmt.Fscan(r, &nodes[i].cat)

	}
	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(r, &x, &y)
		nodes[x-1].neighbors = append(nodes[x-1].neighbors, &nodes[y-1])
		nodes[y-1].neighbors = append(nodes[y-1].neighbors, &nodes[x-1])

	}
	return &nodes[0], m
}
func Kefa(root *node, m int) int {
	var kefa func(cur, prev *node, cnt int, w, m int) int
	kefa = func(cur, prev *node, cnt int, w, m int) int {
		if cur.cat == 0 {
			cnt = 0
		} else {
			cnt += cur.cat
			if cnt > m {
				return w
			}
		}
		if len(cur.neighbors) == 1 && prev != nil {
			return w + 1
		}
		for _, v := range cur.neighbors {
			if v != prev {
				w = kefa(v, cur, cnt, w, m)
			}
		}
		return w
	}
	return kefa(root, nil, 0, 0, m)

}
func main() {
	fmt.Println(Kefa(data(bufio.NewReader(os.Stdin))))
}
