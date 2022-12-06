//https://codeforces.com/contest/580/submission/184129455

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
	Parsing:
	take n and m.  n - node amount  m - resistance of Kefa
	parse line with elements is there cat or not
	parse all following lines with connections of the nodes

	Implementation:
	iterate through tree from root to leaves recurrsively
	Check if there is a cat in node and count it
	When k > m go back
	When we went through all tree, return visited leaves number
*/

type node struct {
	hasCat      int
	connections []*node
}

//func n(hasCat int) *node { return &node{hasCat: hasCat} }
func (n *node) c(c *node) *node {
	n.connections = append(n.connections, c)
	c.connections = append(c.connections, n)
	return n
}

func Kefa(t *node, m int) int {
	var restaurants int
	var kefa func(t *node, p *node, cats int)
	kefa = func(t *node, p *node, cats int) {
		if t == nil {
			return
		}
		if t.hasCat == 1 {
			cats++
		} else {
			cats = 0
		}
		if cats > m { // if kefa can't bare cats anymore
			return
		}
		if len(t.connections) == 1 && t.connections[0] == p { // if we have only connection with parent
			restaurants++ //it's leaf node
		}

		for _, connection := range t.connections {
			if connection != p { // if connection is not parent for us not to return back
				kefa(connection, t, cats)
			}
		}
	}
	kefa(t, nil, 0)
	return restaurants

}

func main() {
	//txt, _ := os.Open("kefa.txt")
	file := bufio.NewReader(os.Stdin /*txt*/)
	fmt.Println(Kefa(input(file)))
}

// func tree() *node {
// 	return n(1).c(n(0)).c(n(0)).c(n(1))
// 	//return n(1).c(n(0).c(n(1)).c(n(0))).c(n(1).c(n(0)).c(n(0)))
// }

func input(file io.Reader) (*node, int) {

	//taking first 2 values
	var n, m int //n - node amount  m - resistance of Kefa
	fmt.Fscan(file, &n, &m)

	//parsing line with cats
	t := make([]*node, n) //creating tree with n elements
	for i := range t {
		var cat int
		fmt.Fscan(file, &cat)     // reading each cat
		t[i] = &node{hasCat: cat} // assigning new node with cat bool to each node link in tree
	}

	//parsing lines with node connections
	for i := 1; i < n; i++ {
		var x, y int
		fmt.Fscan(file, &x, &y)
		x--
		y--
		t[x].c(t[y])
	}
	return t[0], m
}
