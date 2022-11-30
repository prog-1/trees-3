package main

import "fmt"

type node struct {
	v           int
	right, left *node
}

func find(root *node, m map[int]int) map[int]int {
	if root == nil {
		return m
	}
	find(root.left, m)
	m[root.v]++
	find(root.right, m)
	return m
}

func checkduplications(m map[int]int) bool {
	fmt.Println(m)
	for i := 0; i <= len(m)-1; i++ {
		fmt.Print(m[i])
		if m[i] > 1 {
			return true
		}
	}
	return false
}

func main() {
	m := make(map[int]int)
	root := &node{v: 5, left: &node{v: 9, right: &node{v: 10}}, right: &node{v: 10, right: &node{v: 1, left: &node{v: 2}}}}
	if checkduplications(find(root, m)) {
		fmt.Println("There is duplications in tree")
	} else {
		fmt.Println("No duplications in tree")
	}
}
