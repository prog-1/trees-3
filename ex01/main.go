package main

//https://codeforces.com/contest/580/submission/183400304
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Tree struct {
	value    int
	children []*Tree
}

func (tree *Tree) Visit(parent *Tree, maxCats int, catCount int, res *int) {
	catCount += tree.value
	if catCount > maxCats {
		return
	}
	if len(tree.children) == 1 && parent != nil {
		*res += 1
		return
	}
	for _, child := range tree.children {
		if child == parent {
			continue
		}
		child.Visit(tree, maxCats, catCount*child.value, res)
	}
}

func main() {
	var vertices, maxCats int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	fmt.Scanf("%v %v", &vertices, &maxCats)

	tree := make([]Tree, vertices+1, vertices+1)

	for i := 0; i < vertices; i++ {
		scanner.Scan()
		tree[i+1].value, _ = strconv.Atoi(scanner.Text())
	}
	for i := 0; i < vertices-1; i++ {
		scanner.Scan()
		from, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		to, _ := strconv.Atoi(scanner.Text())
		tree[from].children = append(tree[from].children, &tree[to])
		tree[to].children = append(tree[to].children, &tree[from])
	}
	res := 0
	tree[1].Visit(nil, maxCats, 0, &res)
	fmt.Println(res)
}
